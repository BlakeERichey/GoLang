package main

import (
	network "GoLang/NN"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	//load input, targets from wines
	inputs, targets := load("csvs")
	fmt.Println("Loaded", len(inputs), "days")

	//build network
	nn := network.Sequential()
	nn.AddLayer(50, "input")  //add input layer
	nn.AddLayer(256, "relu")  //add hidden layer
	nn.AddLayer(256, "relu")  //add hidden layer
	nn.AddLayer(256, "relu")  //add hidden layer
	nn.AddLayer(2, "softmax") //add output  layer
	nn.Compile(true)
	nn.Summary() //Print Network Params

	//Configure NNEvo Params
	config := network.Config{
		Population:  100,
		Generations: 10000,
		Elites:      30,
		Goal:        .9,
		Metric:      "valid-acc",
	}
	agents := network.NewNNEvo(&config) //initialize with params
	agents.CreatePopulation(nn)         //create network pool
	agents.Summary()

	//run algorithm and return best model after config.Generations
	start := time.Now()
	feed, exp, valFeed, valExp := network.ValidationSplit(inputs, targets, .3)
	model := agents.Fit(
		feed,
		exp,
		valFeed,
		valExp,
		"cross-entropy",
		50,
	)
	elapsed := (time.Now()).Sub(start)
	fmt.Println("Elapsed Time:", elapsed)

	//View models results
	loss, acc := model.Evaluate(inputs, targets, "cross-entropy")
	fmt.Println("Loss:", loss, "Acc:", acc)
	model.Save("stocks.model")
}

func load(dir string) ([][]float64, [][]float64) {
	var files []string
	unableToLoadFiles := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if dir != path {
			files = append(files, path)
		}
		return nil
	})
	if unableToLoadFiles != nil {
		panic(unableToLoadFiles)
	}

	if len(files) < 50 {
		for _, file := range files {
			fmt.Println("Loaded", file)
		}
	}

	cummCsv := make([][]string, 0)
	for _, filename := range files {
		data, _ := ioutil.ReadFile(filename)
		rec := csv.NewReader(strings.NewReader(string(data)))
		records, err := rec.ReadAll()
		if err != nil {
			log.Fatal(err)
		}
		cummCsv = append(cummCsv, records[1:]...)
	}

	ids := make([]float64, len(cummCsv))
	inputs := make([][]float64, 0)
	// outputs := make([][]float64, 0)
	for i := range cummCsv {
		//row values as floats
		lstm := make([]float64, 50)
		data := make([]float64, 50)

		//convert values from string
		lstmarr := strings.Split(cummCsv[i][1][1:len(cummCsv[i][1])-1], ", ")
		dataarr := strings.Split(cummCsv[i][2][1:len(cummCsv[i][2])-1], ", ")
		for j, val := range lstmarr {
			//get lstm value
			datapoint, err := strconv.ParseFloat(val, 64)
			if err != nil {
				log.Fatal(err)
			}
			lstm[j] = datapoint

			//get data value
			datapoint, err = strconv.ParseFloat(dataarr[j], 64)
			if err != nil {
				log.Fatal(err)
			}
			data[j] = datapoint
		}
		//get closing value
		closing, err := strconv.ParseFloat(cummCsv[i][3], 64)
		if err != nil {
			log.Fatal(err)
		}
		if data[len(data)-1] < closing {
			ids[i] = 0
		} else {
			ids[i] = 1
		}
		inputs = append(inputs, lstm)
	}
	targets := onehot(2, ids)

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
