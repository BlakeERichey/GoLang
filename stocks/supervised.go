package main

import (
	network "GoLang/NN"
	"fmt"
	"math/rand"
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

	nn = network.Load("stocks.model")
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
	model.Save("supervised.model")
}
