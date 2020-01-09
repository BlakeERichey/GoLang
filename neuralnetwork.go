package main

import (
	network "GoLang/NN"
	"fmt"
)

func main() {
	nn := network.Sequential()
	nn.AddLayer(5, "input")
	nn.AddLayer(3, "linear")
	nn.AddLayer(5, "linear")
	nn.AddLayer(7, "linear")
	nn.AddLayer(3, "softmax")
	nn.Compile()

	data := [][]float64{
		{2.0, 1.0, 3.0, 4.0, 5.0},
		{-3.7, 4.7, -5.7, -6.7, 7.7},
		{14, 0.123, 0.89, -6.7, 7.0},
	}

	fmt.Println(nn, "\n\n")
	outputs := nn.FeedFoward(data)

	fmt.Println(outputs)

	nn.Save("test.json")
	nn = network.Load("test.json")
}
