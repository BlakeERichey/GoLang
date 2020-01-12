package main

import (
	network "GoLang/NN"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	//build network
	nn := network.Sequential()
	nn.AddLayer(50, "input") //add input layer
	nn.AddLayer(256, "relu") //add hidden layer
	nn.AddLayer(256, "relu") //add hidden layer
	nn.AddLayer(256, "relu") //add hidden layer
	nn.AddLayer(1, "tanh")   //add output  layer
	nn.Compile(true)
	nn.Summary()

	//create Env
	env := InitStockEnv()

	//Configure NNEvo Params
	config := network.Config{
		Population:  100,
		Generations: 1000,
		Elites:      20,
		Goal:        10,
		Metric:      "reward",
	}
	agents := network.NewNNEvo(&config) //initialize with params
	agents.CreatePopulation(nn)         //create network pool
	agents.NewContEnv(env)
	agents.Summary()

	verbosity := 1
	sharpness := 1
	validate := true
	start := time.Now()
	nn = agents.Train(validate, sharpness, verbosity)
	elapsed := (time.Now()).Sub(start)
	fmt.Println("Elapsed Time:", elapsed)
	nn.Save("ga.model")

	reward, valid := network.RunCont(env, nn, sharpness, validate, true)
	fmt.Println("Reward:", reward, "Validation:", valid)
}
