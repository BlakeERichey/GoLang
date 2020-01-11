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
	validMetrics = []string{"acc", "loss", "valid-loss", "valid-acc"}
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

	//used for mutations
	shapes    [][]int
	layerLocs []int //last index in weights for each layer
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
	if config.Population < config.Elites {
		panic("Incapable of taking more elites than entire population.")
	}
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

//Summary prints out a friend display of the NNEvo Layout
func (agents *NNEvo) Summary() {
	fmt.Println(
		"\n==================================NNEvo "+
			"Summary=================================",
		"\nGenerations:", agents.generations,
		"\nPopulaiton:", len(agents.population),
		"\nElites:", agents.elites,
		"\nGoal:", agents.goal,
		"\nMetric:", agents.metric,
		"\nMxrt:", agents.mxrt,
		"\n========================================"+
			"========================================")
}

//CreatePopulation creates duplicates of provided network equal to Config.Population
//nn: reference neural network
func (agents *NNEvo) CreatePopulation(nn *Network) {
	for i := range agents.population {
		agents.population[i] = nn.Clone()
	}
	shapes, weights, _, _ := nn.Serialize()
	agents.shapes = shapes
	agents.layerLocs = make([]int, len(shapes))
	for i, val := range shapes {
		rows, cols := val[0], val[1]
		if i == 0 {
			agents.layerLocs[i] = rows * cols
		} else {
			agents.layerLocs[i] = rows*cols + agents.layerLocs[i-1]
		}
	}
	if agents.mxrt == 0.0 {
		agents.autoMxrt(len(shapes), weights...)
	}
}

//Fit attempts to generate a neural net that produces targets when inputs
//are fed into the network
//method: loss function (i.e. "mse", "cross-entropy")
//verbosity: generations to go without logging results
func (agents *NNEvo) Fit(inputs, targets, validInputs, validTargets [][]float64, method string, verbosity int) *Network {
	if !contains(validLosses, method) {
		panic("Invalid loss type.")
	}
	if !(len(targets) > 0) {
		panic("Unable to compute loss with no targets.")
	}
	if (validInputs == nil || validTargets == nil) && contains([]string{"valid-loss", "valid-acc"}, agents.metric) {
		panic("No validation data to evaluate validation metric")
	}
	// fmt.Println("Fit params:", len(inputs), len(targets), len(validInputs), len(validTargets))

	goalMet := false
	var bestModel *Network
	useBias := agents.population[0].bias
	for gen := 0; gen < agents.generations; gen++ {
		losses := make([]float64, len(agents.population))
		for i := range agents.population { //fitness calculation
			outputs := agents.population[i].FeedForward(inputs)
			loss := calcLoss(outputs, targets, method)
			losses[i] = loss
		}
		matingPool := agents.nextGen(losses, true)
		bestModel = agents.population[matingPool[0]]
		var loss, acc, valLoss, valAcc float64
		loss, acc = bestModel.Evaluate(inputs, targets, method)
		// fmt.Println("loss", loss, "acc:", acc)
		if validInputs != nil && validTargets != nil {
			valLoss, valAcc = bestModel.Evaluate(validInputs, validTargets, method)
			// fmt.Println("valLoss", valLoss, "valAcc:", valAcc)
		}
		goalMet = agents.logMetrics(loss, acc, valLoss, valAcc, gen, verbosity)
		if gen != agents.generations-1 && !goalMet {
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
		if goalMet {
			break
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
				//get layer
				var layer int
				for _, val := range agents.layerLocs {
					if j > val {
						layer++
					} else {
						break
					}
				}
				//get mutation limit
				rows, cols := agents.shapes[layer][0], agents.shapes[layer][1]
				limit := math.Sqrt(6.0 / float64((rows + cols)))
				//mutate
				newWeights[i][j] = -limit + rand.Float64()*(2*limit) //glorot uniform
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

func (agents *NNEvo) autoMxrt(numLayers int, weights ...float64) {
	// agents.mxrt = 1.0 / math.Log(float64(len(weights))) / (float64(len(agents.population)))
	// agents.mxrt = 1.0 / math.Log10(float64(len(weights))) / (float64(len(agents.population))) / float64(numLayers)
	// agents.mxrt = float64(numLayers) / float64(len(weights)) / (float64(len(agents.population)))
	agents.mxrt = math.Log(float64(len(weights))) / float64(len(weights)) / float64(numLayers) / math.Log10(float64(len(agents.population)))
}

func (agents *NNEvo) logMetrics(loss, acc, valLoss, valAcc float64, gen, verbosity int) bool {
	goalMet := false
	if verbosity > 0 && gen%verbosity == 0 {
		fmt.Print("Gen " + strconv.Itoa(gen) + ": loss - ")
		metric := strconv.FormatFloat(loss, 'f', 6, 64)
		if agents.metric == "acc" || agents.metric == "valid-acc" {
			metric = metric + " acc - " + strconv.FormatFloat(acc, 'f', 6, 64)
		}
		if agents.metric == "valid-loss" || agents.metric == "valid-acc" {
			metric = metric + " valid-loss - " + strconv.FormatFloat(valLoss, 'f', 6, 64)
			if agents.metric == "valid-acc" {
				metric = metric + " valid-acc - " + strconv.FormatFloat(valAcc, 'f', 6, 64)
			}
		}
		fmt.Println(metric)
	}
	if agents.metric == "loss" && loss < agents.goal {
		goalMet = true
	}
	if agents.metric == "acc" && acc > agents.goal {
		goalMet = true
	}
	if agents.metric == "valid-loss" && valLoss < agents.goal {
		goalMet = true
	}
	if agents.metric == "valid-acc" && valAcc > agents.goal {
		goalMet = true
	}
	return goalMet
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

//ValidationSplit returns input, targets, validInputs, validTargets
//splitting with perc% of the values being present in the validation arrays
func ValidationSplit(inputs, targets [][]float64, perc float64) (in, tar, vI, vT [][]float64) {
	validStart := int((1 - perc) * float64(len(inputs)))
	return inputs[:validStart], targets[:validStart], inputs[validStart:], targets[validStart:]
}
