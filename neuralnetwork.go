package main

import (
	network "GoLang/NN"
	"fmt"
)

func main() {
	nn := network.Sequential(5, 1)
	nn.AddLayer(3, "tanh")
	nn.AddLayer(5, "relu")
	nn.AddLayer(7, "linear")
	nn.Compile("sigmoid")

	data := [][]float64{
		{2.0, 1.0, 3.0, 4.0, 5.0},
		{-3.7, 4.7, -5.7, -6.7, 7.7},
	}

	fmt.Println(nn, "\n\n")
	outputs := nn.FeedFoward(data)

	fmt.Println(outputs)

}
