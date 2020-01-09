package main

import (
	network "GoLang/NN"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	nn := network.Sequential()
	nn.AddLayer(5, "input")
	nn.AddLayer(3, "linear")
	nn.AddLayer(5, "linear")
	nn.AddLayer(7, "linear")
	nn.AddLayer(3, "softmax")
	nn.Compile()
	fmt.Println(nn, "\n\n")

	data := [][]float64{
		{2.0, 1.0, 3.0, 4.0, 5.0},
		{-3.7, 4.7, -5.7, -6.7, 7.7},
		{14, 0.123, 0.89, -6.7, 7.0},
	}

	outputs := nn.FeedFoward(data)
	fmt.Println(outputs)

	nn.Save("model.json")
	nn = network.Load("model.json")

	//log time taken for 100000 predictions: 256ms
	var newData [2250][50]float64
	for i := 0; i < 2250; i++ {
		for j := 0; j < 50; j++ {
			newData[i][j] = rand.Float64() * float64(rand.Intn(50))
		}
	}
	fmt.Println("New data:", newData[0], "...")
	start := time.Now()
	for i := 0; i < 100000; i++ {
		outputs = nn.FeedFoward(data)
	}
	elapsed := (time.Now()).Sub(start)
	fmt.Println("Elapsed Time:", elapsed)
}
