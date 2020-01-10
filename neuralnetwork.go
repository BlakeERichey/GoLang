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
	nn.AddLayer(9, "linear")
	nn.AddLayer(5, "linear")
	nn.AddLayer(1, "tanh")
	nn.Compile(true)
	// fmt.Println(nn, "\n\n")

	config := network.Config{
		Population:  100,
		Generations: 100,
		Elites:      1,
		Goal:        .99,
		Metric:      "acc",
		Mxrt:        0.01,
	}
	agents := network.NewNNEvo(&config)
	agents.CreatePopulation(nn)

	inputs := [][]float64{
		{2.0, 1.0, 3.0, 4.0, 5.0},
		{-3.7, 4.7, -5.7, -6.7, 7.7},
		{14, 0.123, 0.89, -6.7, 7.0},
	}
	targets := [][]float64{
		{0},
		{1},
		{1},
	}

	model := agents.Fit(inputs, targets, "mse", 1)

	preds := model.FeedFoward(inputs)
	fmt.Println(preds)
	model.Save("model.json")

	// outputs := nn.FeedFoward(inputs)
	// fmt.Println(outputs)

	// shapes, weights, _, bias := nn.Serialize()
	// newWeights := make([]float64, len(weights))
	// for i := range bias {
	// 	for j := range bias[i] {
	// 		bias[i][j] = 0
	// 	}
	// }
	// des := network.Deserialize(shapes, newWeights)
	// nn.SetWeights(des...)
	// nn.SetBias(bias...)
	// fmt.Println(nn, "\n\n")
	// nn.Save("model.json")
	// nn = network.Load("model.json")

	// nn := network.Sequential()
	// nn.AddLayer(50, "input")
	// nn.AddLayer(256, "relu")
	// nn.AddLayer(256, "relu")
	// nn.AddLayer(256, "relu")
	// nn.AddLayer(1, "tanh")
	// nn.Compile(false)
	// nn.Summary()

	// //log time taken for 1000 predictions: 28.7s
	// newData := make([][]float64, 0)
	// for i := 0; i < 2250; i++ {
	// 	ob := make([]float64, 0)
	// 	for j := 0; j < 50; j++ {
	// 		ob = append(ob, rand.Float64()*float64(rand.Intn(50)))
	// 	}
	// 	newData = append(newData, ob)
	// }
	// fmt.Println("New data:(", len(newData[0]), ",", len(newData), ")...")
	// start := time.Now()
	// for i := 0; i < 1000; i++ {
	// 	shapes, weights, _, _ := nn.Serialize()
	// 	des := network.Deserialize(shapes, weights)
	// 	nn.SetWeights(des...)
	// 	nn.FeedFoward(newData)
	// }
	// elapsed := (time.Now()).Sub(start)
	// fmt.Println("Elapsed Time:", elapsed)
}
