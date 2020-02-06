package main

import (
	network "GoLang/NN"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	//load input, targets from wines
	inputs, targets := load("iris.data")
	// fmt.Println(inputs, targets)
	inputs_train, targets_train, validInputs, validTargets := network.ValidationSplit(inputs, targets, .2)

	//build network
	nn := network.Sequential()
	nn.AddLayer(4, "input")  //add input layer
	nn.AddLayer(8, "relu")   //add hidden layer
	nn.AddLayer(12, "relu")   //add hidden layer
	nn.AddLayer(8, "relu")   //add hidden layer
	nn.AddLayer(3, "softmax") //add output  layer
	nn.Compile(false)
	nn.Summary() //Print Network Params

	//Configure NNEvo Params
	config := network.Config{
		Population:  100,
		Generations: 10000,
		Elites:      30,
		Goal:        .995,
		Metric:      "valid-acc",
		// Mxrt:        0.001, //Do not include mxrt for NNEvo to autoinfer a mutation rt
	}
	agents := network.NewNNEvo(&config) //initialize with params
	agents.CreatePopulation(nn)         //create network pool
	agents.Summary()

	//run algorithm and return best model after config.Generations
	start := time.Now()
	model := agents.Fit(inputs_train, targets_train, validInputs, validTargets, "cross-entropy", 100)
	elapsed := (time.Now()).Sub(start)
	fmt.Println("Elapsed Time:", elapsed)

	//View models results
	loss, acc := model.Evaluate(validInputs, validTargets, "cross-entropy")
	fmt.Println("Loss:", loss, "Acc:", acc)
	model.Save("iris.model")
}

func load(filename string) ([][]float64, [][]float64) {
	data, _ := ioutil.ReadFile(filename)
	rec := csv.NewReader(strings.NewReader(string(data)))
	records, err := rec.ReadAll()
	if err != nil {
		panic("Unable to parse file.")
	}

	ids := make([]float64, len(records))
	inputs := make([][]float64, 0)
	fmt.Println("Records:", len(records), "Attributes:", len(records[0])-1)
	for i := range records {
		val, err := strconv.ParseFloat(records[i][len(records[i])-1], 64)
		if err == nil {
			ids[i] = val
		}
		input := make([]float64, len(records[0])-1)
		for j := range input {
			val, err2 := strconv.ParseFloat(records[i][j], 64)
			if err2 == nil {
				input[j] = val
			}
		}
		inputs = append(inputs, input)
	}
	

	targets := onehot(3, ids)
	return inputs, targets
}

func onehot(classes int, data []float64) [][]float64 {
	encoded := make([][]float64, 0)
	for _, val := range data {
		categorical := make([]float64, classes)
		categorical[int(val)] = 1.0
		encoded = append(encoded, categorical)
	}
	return encoded
}