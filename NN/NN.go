package network

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"os"

	"gonum.org/v1/gonum/blas/blas64"
	"gonum.org/v1/gonum/mat"
)

var (
	validActivations = []string{"relu", "linear", "softmax", "sigmoid", "tanh"}
)

type Network struct {
	layers []Layer
	input  int //number of input nodes
	output int //number of output nodes
}

//Sequential is used to build a Network one layer at a time
func Sequential() *Network {
	return new(Network)
}

//AddLayer adds a Dense layer to the end of the Network
func (nn *Network) AddLayer(nodes int, activation string) {
	var rows int
	if nn.input == 0 {
		if activation != "input" && len(activation) > 0 {
			panic("Input layer should not have an attached activation")
		}
		nn.input = nodes
	} else if len(nn.layers) == 0 {
		rows = nn.input
	} else {
		rows = nn.layers[len(nn.layers)-1].cols
	}

	if rows > 0 { //ignore first layer, as is insuffient information to gen weights
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
}

//Compile closes Network and notates the last layer.
func (nn *Network) Compile() { //needs LR & Optimizer later
	nn.output = nn.layers[len(nn.layers)-1].weights.RawMatrix().Cols
}

//FeedForward passes data into the Network and generates an output.
//FeedForward is capable of handling multiple inputs at once.
func (nn *Network) FeedFoward(data [][]float64) [][]float64 {
	obs := len(data)
	inputs := make([]float64, 0)
	predictions := make([][]float64, obs)
	for i := 0; i < obs; i++ {
		if len(data[i]) != nn.input {
			fmt.Println("Expected input shape", nn.input, ", recevied", len(data[i]))
			panic("Invalid input structure")
		}
		inputs = append(inputs, data[i][:nn.input]...)
	}
	var hi mat.Dense //hidden layer input matrix for next layer
	hi = *mat.NewDense(obs, nn.input, inputs)
	for j := range nn.layers {
		var ho mat.Dense //hidden layer output matrix
		// fmt.Println("HI:", hi)
		// fmt.Println("HW:", nn.layers[j].weights)
		ho.Mul(&hi, &nn.layers[j].weights)
		//activations happen here!!!
		applyActivation(&nn.layers[j], ho.RawMatrix(), ho.RawMatrix().Data...)
		hi = ho
	}
	start := 0                     //where next prediction starts in array
	outputs := hi.RawMatrix().Data //resulting matrix
	for i := range predictions {
		predictions[i] = outputs[start : nn.output+start]
		start += nn.output
	}
	return predictions
}

//applyActivation uses layers activation method to modify the
//hidden layer output matrix in place
func applyActivation(layer *Layer, matrix blas64.General, data ...float64) {
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
	} else if layer.activation == "softmax" {
		stride := 0
		exp := make([]float64, layer.cols)
		for ob := 0; ob < matrix.Rows; ob++ {
			for i := range exp { //raise e^data[i]
				exp[i] = math.Exp(data[stride+i])
			}
			//get the sum of the new array
			total := 0.0
			for _, val := range exp {
				total += val
			}
			//gen new values
			for i := range exp {
				data[stride+i] = exp[i] / total
			}
			stride += layer.cols
		}
	}
}

//Serialize returns relevant data to regenerate the Network
func (nn *Network) Serialize() (shapes [][]int, weights []float64, activations []string) {
	shapes = make([][]int, 0) //[[y, x], ...]
	weights = make([]float64, 0)
	activations = make([]string, 0)

	for i := range nn.layers {
		raw := nn.layers[i].weights.RawMatrix()
		shapes = append(shapes, []int{raw.Rows, raw.Cols})
		weights = append(weights, raw.Data...)
		activations = append(activations, nn.layers[i].activation)
	}
	// fmt.Println("Shapes:", shapes)
	// fmt.Println("Weights:", weights)
	return shapes, weights, activations
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

type Model struct {
	Name        string
	Shapes      [][]int
	Weights     []float64
	Activations []string
}

//Save serializes and saves a network into a json
func (nn *Network) Save(filename string) {
	shapes, weights, activations := nn.Serialize()
	data := Model{
		Name:        filename,
		Shapes:      shapes,
		Activations: activations,
		Weights:     weights,
	}
	file, _ := json.Marshal(data)
	err := ioutil.WriteFile(data.Name, file, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

//Load takes a filename and returns a Network. Expects a json generated via
//nn.Save
func Load(filename string) *Network {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, _ := ioutil.ReadAll(file)
	var model Model
	err = json.Unmarshal(data, &model)

	//create Network
	var nn Network
	input, output := model.Shapes[0][0], model.Shapes[len(model.Shapes)-1][1]
	nn.input, nn.output = input, output
	curIndex := 0 //current starting weight index for layer
	for i := range model.Shapes {
		rows, cols := model.Shapes[i][0], model.Shapes[i][1]
		nn.layers = append(nn.layers, *createLayer(rows, cols, model.Activations[i], model.Weights[curIndex:rows*cols+curIndex]))
		curIndex += rows * cols
	}
	return &nn
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
