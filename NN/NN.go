package network

import (
	"fmt"
	"math"
	"math/rand"

	"gonum.org/v1/gonum/mat"
)

var (
	validActivations = []string{"relu", "linear", "sigmoid", "tanh"}
)

type Network struct {
	layers []Layer
	input  int
	output int
}

func Sequential(input, output int) *Network {
	var network Network
	network.input = input
	network.output = output
	return &network
}

func (nn *Network) AddLayer(nodes int, activation string) {
	var rows int
	if len(nn.layers) == 0 {
		rows = nn.input
	} else {
		rows = nn.layers[len(nn.layers)-1].cols
	}
	//initialize weights
	newWeights := make([]float64, rows*nodes)
	limit := math.Sqrt(6.0 / float64((rows + nodes)))
	for j := range newWeights {
		newWeights[j] = -limit + rand.Float64()*(2*limit)
	}
	//create layer
	layer := *createLayer(rows, nodes, activation, newWeights)
	//add layer to Network
	nn.layers = append(nn.layers, layer)
}

func (nn *Network) Compile(activation string) { //needs LR & Optimizer later
	nn.AddLayer(nn.output, activation)
}

func (nn *Network) FeedFoward(data [][]float64) [][]float64 {
	obs := len(data)
	outputs := make([][]float64, obs)
	for i := 0; i < obs; i++ {
		if len(data[i]) != nn.input {
			fmt.Println("Expected input shape", nn.input, ", recevied", len(data[i]))
			panic("Invalid input structure")
		}
		input := data[i][:nn.input]
		var hi mat.Dense //hidden layer input matrix for next layer
		hi = *mat.NewDense(1, nn.input, input)
		for j := range nn.layers {
			var ho mat.Dense //hidden layer output matrix
			// fmt.Println("HI:", hi)
			// fmt.Println("HW:", nn.layers[j].weights)
			ho.Mul(&hi, &nn.layers[j].weights)
			//activations happen here!!!
			applyActivation(&nn.layers[j], ho.RawMatrix().Data...)
			hi = ho
		}
		outputs[i] = hi.RawMatrix().Data
	}
	return outputs
}

func applyActivation(layer *Layer, data ...float64) {
	if !Contains(validActivations, layer.activation) {
		panic("Invalid Activation")
	}
	if layer.activation == "relu" {
		for i := range data {
			data[i] = math.Max(data[i], 0)
		}
	} else if layer.activation == "sigmoid" {
		for i := range data {
			data[i] = 1.0 / (1 + math.Exp(-data[i]))
		}
	} else if layer.activation == "tanh" {
		for i := range data {
			data[i] = math.Tanh(data[i])
		}
	} else if layer.activation == "sigmoid" {
		for i := range data {
			data[i] = 1.0 / (1 + math.Exp(-data[i]))
		}
	}
}

func (nn *Network) Serialize() ([][]int, []float64) {
	shapes := make([][]int, 0) //[[y, x], ...]
	weights := make([]float64, 0)

	for i := range nn.layers {
		raw := nn.layers[i].weights.RawMatrix()
		shapes = append(shapes, []int{raw.Rows, raw.Cols})
		weights = append(weights, raw.Data...)
	}
	// fmt.Println("Shapes:", shapes)
	// fmt.Println("Weights:", weights)
	return shapes, weights
}

//Deserialize takes in the shapes and weights of a network and then returns
//the Layer.weights values back as a list. This can be used to set layer weights
//for an entire network
func Deserialize(shapes [][]int, weights []float64) (desMat []mat.Dense) {
	desMat = make([]mat.Dense, len(shapes))
	curIndx := 0
	for i := range shapes {
		rows, cols := shapes[i][0], shapes[i][1]
		end := rows*cols + curIndx
		fmt.Println(curIndx, end)
		desMat[i] = *mat.NewDense(rows, cols, weights[curIndx:end])
		curIndx = end
	}
	return desMat
}

type Layer struct {
	rows       int
	cols       int
	weights    mat.Dense
	activation string
	//bias []float32
}

func createLayer(rows, cols int, activation string, data []float64) *Layer {
	layer := *new(Layer)
	layer.cols, layer.rows = cols, rows
	layer.weights = *mat.NewDense(rows, cols, data)
	if !Contains(validActivations, activation) {
		panic("Invalid Activation")
	}
	layer.activation = activation
	return &layer
}

func (layer *Layer) SetWeights(data []float64) {
	layer.weights = *mat.NewDense(layer.rows, layer.cols, data)
}

//Contains returns true is list contains an element == val
//list: list of values to look at
//val:	val to compare elements of list to
func Contains(list []string, val string) bool {
	for _, ele := range list {
		if ele == val {
			return true
		}
	}
	return false
}
