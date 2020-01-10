package network

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"os"
	"reflect"

	"gonum.org/v1/gonum/blas/blas64"
	"gonum.org/v1/gonum/mat"
)

var (
	validActivations = []string{"relu", "linear", "softmax", "sigmoid", "tanh"}
)

type Network struct {
	layers []Layer
	input  int  //number of input nodes
	output int  //number of output nodes
	bias   bool //true if model layers use bias
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

//Compile closes Network and notates the last layer. Adds bias factor to layers
func (nn *Network) Compile(useBias bool) { //needs LR & Optimizer later
	nn.output = nn.layers[len(nn.layers)-1].weights.RawMatrix().Cols
	nn.bias = useBias
	//gen bias
	for i := range nn.layers {
		layer := &nn.layers[i]
		bias := make([]float64, layer.cols)
		var biasVal float64
		if useBias {
			biasVal = -1 + rand.Float64()*(2) //random float [-1,1]
		} else {
			biasVal = 0.0
		}
		for j := range bias {
			bias[j] = biasVal
		}
		layer.bias = bias
	}
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
		if nn.bias {
			applyBias(&nn.layers[j], ho.RawMatrix(), ho.RawMatrix().Data...)
		}
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

//adds layers bias to hidden layer output before activation function
func applyBias(layer *Layer, matrix blas64.General, data ...float64) {
	obStartIndx := 0
	for ob := 0; ob < matrix.Rows; ob++ {
		for i := 0; i < layer.cols; i++ {
			data[obStartIndx+i] += layer.bias[i]
		}
		obStartIndx += layer.cols
	}
}

//applyActivation uses layers activation method to modify the
//hidden layer output matrix in place
func applyActivation(layer *Layer, matrix blas64.General, data ...float64) {
	if !contains(validActivations, layer.activation) {
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

//Summary prints out a friend display of the Network architecture of settings
func (nn *Network) Summary() {
	shapes, _, activations, _ := nn.Serialize()
	fmt.Println(
		"Inputs:", nn.input,
		"\nOutputs:", nn.output,
		"\nNum Layers:", len(nn.layers),
		"\nUses Bias:", nn.bias,
		"\nWeight Shapes:", shapes,
		"\nLayer activations:", activations,
		"\n========================================"+
			"========================================")
}

//Serialize returns relevant data to regenerate the Network
func (nn *Network) Serialize() ([][]int, []float64, []string, [][]float64) {
	bias := make([][]float64, 0)
	shapes := make([][]int, 0) //[[y, x], ...]
	weights := make([]float64, 0)
	activations := make([]string, 0)

	for i := range nn.layers {
		layer := &nn.layers[i]
		raw := layer.weights.RawMatrix()
		shapes = append(shapes, []int{raw.Rows, raw.Cols})
		//prohibits mutation by indexing
		addBias := make([]float64, len(layer.bias))
		if nn.bias {
			copy(addBias, layer.bias) //copy array values instead of using array.
		}
		bias = append(bias, addBias)
		weights = append(weights, raw.Data...)
		activations = append(activations, layer.activation)
	}
	// fmt.Println("Shapes:", shapes)
	// fmt.Println("Weights:", weights)
	return shapes, weights, activations, bias
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
		desMat[i] = *mat.NewDense(rows, cols, weights[curIndx:end])
		curIndx = end
	}
	return desMat
}

//SetWeights takes a array of Dense matrices and applies them to
//Network layers. Expects len(desMat) to be equal to number of Network Layers
func (nn *Network) SetWeights(desMat ...mat.Dense) {
	if len(desMat) != len(nn.layers) {
		panic("Invalid weight structure!")
	}
	for i := range desMat {
		nn.layers[i].weights = desMat[i]
	}
}

func (nn *Network) SetBias(bias ...[]float64) {
	if len(bias) != len(nn.layers) {
		panic("Invalid bias structure!")
	}
	if !nn.bias {
		nn.bias = true
	}
	for i := range bias {
		layer := &nn.layers[i]
		layer.bias = bias[i]
	}
}

type Model struct {
	Name        string
	Shapes      [][]int
	Weights     []float64
	Bias        [][]float64
	Activations []string
}

//Save serializes and saves a network into a json
func (nn *Network) Save(filename string) {
	shapes, weights, activations, bias := nn.Serialize()
	data := Model{
		Name:        filename,
		Shapes:      shapes,
		Activations: activations,
		Weights:     weights,
		Bias:        bias,
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
		layer := *createLayer(rows, cols, model.Activations[i], model.Weights[curIndex:rows*cols+curIndex])
		if !nn.bias {
			emptyBias := make([]float64, layer.cols)
			if !reflect.DeepEqual(emptyBias, model.Bias[i]) {
				nn.bias = true
			}
		}
		layer.bias = model.Bias[i]
		nn.layers = append(nn.layers, layer)
		curIndex += rows * cols
	}
	return &nn
}

//Clone returns a new network with the same architecture but new weights and biases
func (nn *Network) Clone() *Network {
	newNetwork := Sequential()
	inputCreated := false
	i := 0
	for i < len(nn.layers) {
		var nodes int
		var activation string
		layer := &nn.layers[i]
		if i == 0 && !inputCreated {
			nodes = layer.rows
			activation = "input"
			inputCreated = true
		} else {
			nodes = layer.cols
			activation = layer.activation
			i++
		}
		newNetwork.AddLayer(nodes, activation)
	}
	newNetwork.Compile(nn.bias)
	return newNetwork
}

type Layer struct {
	rows       int
	cols       int
	weights    mat.Dense
	bias       []float64
	activation string
}

func createLayer(rows, cols int, activation string, data []float64) *Layer {
	layer := *new(Layer)
	layer.cols, layer.rows = cols, rows
	layer.weights = *mat.NewDense(rows, cols, data)
	if !contains(validActivations, activation) {
		panic("Invalid Activation")
	}
	layer.activation = activation
	return &layer
}

func (layer *Layer) SetWeights(data []float64) {
	layer.weights = *mat.NewDense(layer.rows, layer.cols, data)
}

//contains returns true is list contains an element == val
//list: list of values to look at
//val:	val to compare elements of list to
func contains(list []string, val string) bool {
	for _, ele := range list {
		if ele == val {
			return true
		}
	}
	return false
}
