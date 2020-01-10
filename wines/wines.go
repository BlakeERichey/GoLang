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

	inputs, targets := load("wines.csv")
	nn := network.Sequential()
	nn.AddLayer(13, "input")
	nn.AddLayer(8, "relu")
	nn.AddLayer(8, "relu")
	nn.AddLayer(3, "softmax")
	nn.Compile(true)
	nn.Summary()

	config := network.Config{
		Population:  40,
		Generations: 100000,
		Elites:      10,
		Goal:        .995,
		Metric:      "acc",
		Mxrt:        0.001,
	}
	agents := network.NewNNEvo(&config)
	agents.CreatePopulation(nn)

	start := time.Now()
	model := agents.Fit(inputs, targets, "cross-entropy", 100)
	elapsed := (time.Now()).Sub(start)
	fmt.Println("Elapsed Time:", elapsed)

	//View models results
	loss, acc := model.Evaluate(inputs, targets, "cross-entropy")
	fmt.Println("Loss:", loss, "Acc:", acc)
	model.Save("wines.model")
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
		val, err := strconv.ParseFloat(records[i][0], 64)
		if err == nil {
			ids[i] = val - 1
		}
		input := make([]float64, len(records[0])-1)
		for j := range input {
			val, err2 := strconv.ParseFloat(records[i][j+1], 64)
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
