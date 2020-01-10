package network

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

//fixed topological genetic algorithm neural network weight optimizer

var (
	validLosses  = []string{"mse", "cross-entropy"}
	validMetrics = []string{"acc", "loss"}
)

type NNEvo struct {
	population  []*Network
	generations int
	elites      int
	mxrt        float64
	goal        float64
	metric      string
	validation  bool //used for Env interfaces for stochastic environments
	env         Env
}

type Config struct {
	Population  int
	Generations int
	Elites      int
	Goal        float64
	Metric      string
	Mxrt        float64
}

type Env interface {
	init()
	step() []float64
	reset() []float64
	getSteps() int
	getReward() float64
}

//NewNNEvo Generates a new NNEvo struct from provided Config struct.
//Must call CreatePopulation afterwards to generate neural nets
func NewNNEvo(config *Config) *NNEvo {
	net := new(NNEvo)
	net.goal = config.Goal
	net.mxrt = config.Mxrt
	net.elites = config.Elites
	net.metric = config.Metric
	net.generations = config.Generations
	net.population = make([]*Network, config.Population)
	if !contains(validMetrics, net.metric) {
		panic("Invalid metric.")
	}
	return net
}

//CreatePopulation creates duplicates of provided network equal to Config.Population
//nn: reference neural network
func (agents *NNEvo) CreatePopulation(nn *Network) {
	for i := range agents.population {
		agents.population[i] = nn.Clone()
	}
}

//Fit attempts to generate a neural net that produces targets when inputs
//are fed into the network
//method: loss function
//verbosity: generations to go without logging results
func (agents *NNEvo) Fit(inputs, targets [][]float64, method string, verbosity int) *Network {
	if !contains(validLosses, method) {
		panic("Invalid loss type.")
	}
	if !(len(targets) > 0) {
		panic("Unable to compute loss with no targets.")
	}

	var metric string //value to report
	var bestModel *Network
	useBias := agents.population[0].bias
	for gen := 0; gen < agents.generations; gen++ {
		losses := make([]float64, len(agents.population))
		for i := range agents.population {
			outputs := agents.population[i].FeedFoward(inputs)
			// fmt.Println("output", i, outputs)
			loss := calcLoss(outputs, targets, method)
			// fmt.Println(loss)
			losses[i] = loss
		}
		matingPool := agents.nextGen(losses, true)
		bestModel = agents.population[matingPool[0]]
		if verbosity > 0 && gen%verbosity == 0 {
			fmt.Print("Gen " + strconv.Itoa(gen) + ": loss - ")
			metric = strconv.FormatFloat(losses[matingPool[0]], 'f', 6, 64)
			if agents.metric == "acc" {
				outputs := bestModel.FeedFoward(inputs)
				acc := calcAcc(outputs, targets)
				metric = metric + " acc - " + strconv.FormatFloat(acc, 'f', 6, 64)
			}
			fmt.Println(metric)
		}
		if gen != agents.generations-1 {
			children := agents.crossover(matingPool...)
			agents.mutate(children...)

			//apply new gen to NNEvo.population
			shapes, _, _, bias := bestModel.Serialize()
			for i, weights := range children {
				des := Deserialize(shapes, weights)
				agents.population[i].SetWeights(des...)
				if useBias {
					agents.population[i].SetBias(bias...)
				}
			}
		}
	}
	return bestModel
}

func (agents *NNEvo) crossover(matingPool ...int) (children [][]float64) {
	children = make([][]float64, len(matingPool))
	for i := 0; i < agents.elites; i++ {
		_, weights, _, _ := agents.population[matingPool[i]].Serialize()
		children[i] = weights
	}

	pool := rand.Perm(len(matingPool))
	child := agents.elites //how many children have been bred
	for child < len(matingPool) {
		parent1 := agents.population[matingPool[child]]
		parent2 := agents.population[matingPool[len(pool)-child-1]]
		children[child] = breed(parent1, parent2)
		child++
	}

	return children
}

func breed(parent1, parent2 *Network) []float64 {
	_, weights1, _, _ := parent1.Serialize()
	_, weights2, _, _ := parent2.Serialize()
	n := len(weights1) //genes to process
	newWeights := make([]float64, n)
	i := 0 //processes genes
	for i < n {
		newWeights[i] = weights1[i]
		i++
		if i < n {
			newWeights[i] = weights2[i]
			i++
		}
	}
	return newWeights
}

func (agents *NNEvo) mutate(newWeights ...[]float64) {
	for i := range newWeights {
		for j := range newWeights[i] {
			if rand.Float64() < agents.mxrt {
				//mutate
				newWeights[i][j] = -1 + rand.Float64()*(2) //random float [-1,1]
			}
		}
	}
}

//returns population indices for mating pool in crossover process
func (agents *NNEvo) nextGen(fitness []float64, minimize bool) []int {
	//comparison sort - least to greatest
	compArr := make([]int, len(fitness)) //indices
	for i := 0; i < len(compArr)-1; i++ {
		for j := i + 1; j < len(compArr); j++ {
			if fitness[i] < fitness[j] {
				if minimize {
					compArr[j]++
				} else {
					compArr[i]++
				}
			} else {
				if minimize {
					compArr[i]++
				} else {
					compArr[j]++
				}
			}
		}
	}

	ranked := make([]int, len(compArr)) //routes sorted
	for i := range ranked {
		ranked[compArr[i]] = i
	}

	//select new generation
	newGen := make([]int, len(agents.population))
	i := 0 //next gen index
	for i < agents.elites {
		newGen[i] = ranked[i]
		i++
	}

	remaining := rand.Perm(len(agents.population))
	for i < len(agents.population) {
		newGen[i] = remaining[i-agents.elites]
		i++
	}
	return newGen

}

func calcLoss(preds, actual [][]float64, method string) float64 {
	if len(preds) != len(actual) {
		panic("Inconsistent predictions and labels.")
	}
	var loss float64
	if method == "mse" {
		for i := range preds { //for each obs
			sumSqErr := 0.0
			for j, val := range preds[i] {
				// fmt.Println("Act:", actual[i][j], "Pred", val)
				sumSqErr += math.Pow((actual[i][j] - val), 2)
			}
			loss += sumSqErr / float64(len(preds[i]))
		}
	} else if method == "cross-entropy" {
		for i := range preds {
			sumCxEnt := 0.0
			for j, val := range preds[i] {
				if actual[i][j] == 1.0 {
					sumCxEnt += -math.Log(math.Max(val, 1E-12))
				} else { //actual == 0
					sumCxEnt += -math.Log(math.Max(1-val, 1E-12))
				}
			}
			loss = loss + sumCxEnt/float64(len(preds[i]))
		}
	}
	return loss / float64(len(preds))
}

func calcAcc(preds, actual [][]float64) (acc float64) {
	numCor := 0
	for i, val := range preds { //for each obs
		if argmax(val...) == argmax(actual[i]...) {
			numCor++
		}
	}
	acc = float64(numCor) / float64(len(preds))
	return acc
}

func argmax(arr ...float64) (index int) {
	index = 0
	maxVal := arr[0]
	for i, val := range arr {
		if val > maxVal {
			index = i
			maxVal = val
		}
	}
	return index
}
