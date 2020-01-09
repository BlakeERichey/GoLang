package main

import (
	network "GoLang/NN"
	"fmt"
)

func main() {
	// nn := network.Sequential(5, 1)
	nn := network.Sequential()
	nn.AddLayer(5, "input")
	nn.AddLayer(3, "linear")
	nn.AddLayer(5, "linear")
	nn.AddLayer(7, "linear")
	nn.AddLayer(3, "softmax")
	nn.Compile()

	data := [][]float64{
		{2.0, 1.0, 3.0, 4.0, 5.0},
		// {-3.7, 4.7, -5.7, -6.7, 7.7},
	}

	fmt.Println(nn, "\n\n")
	outputs := nn.FeedFoward(data)

	fmt.Println(outputs)

}
