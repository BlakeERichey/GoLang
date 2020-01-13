package main

import (
	network "GoLang/NN"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	//create Env
	env := InitStockEnv()

	//build network
	nn := network.Sequential()
	nn.AddLayer(50, "input")   //add input layer
	nn.AddLayer(256, "linear") //add hidden layer
	nn.AddLayer(256, "linear") //add hidden layer
	nn.AddLayer(256, "linear") //add hidden layer
	nn.AddLayer(1, "tanh")     //add output  layer
	nn.Compile(false)
	nn.Summary()

	//Configure NNEvo Params
	config := network.Config{
		Population:  100,
		Generations: 50000,
		Elites:      20,
		Goal:        100,
		Metric:      "reward",
		Callbacks:   []string{"checkpoint"},
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
	nn.Save("reinforce.model")

	reward, valid := network.RunCont(env, nn, sharpness, validate, true)
	fmt.Println("Reward:", reward, "Validation:", valid)
}
