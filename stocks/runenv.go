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
	nn := network.Load("checkpoint.model")
	nn.Summary()
	reward, valid := network.RunCont(env, nn, 1, true, true)
	fmt.Println("Reward:", reward, "Validation:", valid)
	// inputs := env.Reset()
	// fmt.Println("Inputs:", inputs[0])
}
