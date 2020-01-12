# Overview  
This is a Fixed-Topological Neural Network package that uses Genetic Algorithms 
to optimize its weights.  

## Supported Features:  
* Activations: ReLU, Linear, Softmax, Tanh, Sigmoid  
* Layers: Dense  
* Loss functions: Cross-Entropy, Mean-Squared Error

# Installation 
Package location subject to change, so installation instructions will come later.  

# Usage  

## Building a Neural Network  
This package offers the ability to contruct a network one layer at a time.  
To initialize a Neural Network one might run:  
```golang
  package main

  import (
    network "github.com/NN..." //import state for package 
    //other imports
  )

  func main(){
    //build network
    nn := network.Sequential()
    nn.AddLayer(5, "input")   //add input layer
    nn.AddLayer(8, "linear")  //add hidden layer
    nn.AddLayer(8, "relu")    //add hidden layer
    nn.AddLayer(2, "softmax") //add output  layer
    nn.Compile(true)
    nn.Summary() //Print Network Params
  }
```  

Once a Network has been compiled or loaded, one can get predictions 
from the model using the provided FeedForward function:  
```golang  
  package main

  import (
    network "github.com/NN..." //import state for package 
    //other imports
  )

  func main(){
    //build network
    nn := network.Sequential()
    nn.AddLayer(5, "input")   //add input layer
    nn.AddLayer(8, "linear")  //add hidden layer
    nn.AddLayer(8, "relu")    //add hidden layer
    nn.AddLayer(2, "softmax") //add output  layer
    nn.Compile(true)          //compile and use layer bias
    nn.Summary() //Print Network Params

    //Load model instead
    nn = network.Load("example.model")

    inputs := [][]float64{
      {2.0, 1.0, 3.0, 4.0, 5.0},
      {-3.7, 4.7, -5.7, -6.7, 7.7},
      {14, 0.123, 0.89, -6.7, 7.0},
    }

    //get outputs
    outputs := nn.FeedForward(inputs)
    
    //print network outputs
    for _, val := range outputs {
      fmt.Println(val)
    }
    
  }
```  

## Supervised Learning  
This package uses a Genetic Algorithm to optimize weights, so to Fit data 
you need to create a pool of agents:  
```golang
  package main

  import (
    network "github.com/NN..." //import state for package 
    //other imports
  )

  func main(){
    //build network
    nn := network.Sequential()
    nn.AddLayer(5, "input")   //add input layer
    nn.AddLayer(8, "linear")  //add hidden layer
    nn.AddLayer(8, "relu")    //add hidden layer
    nn.AddLayer(2, "softmax") //add output  layer
    nn.Compile(true)          //compile and use layer bias
    nn.Summary() //Print Network Params

    //Build Pool of workers
    //Configure NNEvo Params
    config := network.Config{
      Population:  100,  //how many networks to make
      Generations: 10000,
      Elites:      30,   //the top 30 networks transition to the net generation
      Goal:        .995, //Algorithm stops if this value is reach on metric provided
      Metric:      "acc",
      // Mxrt:        0.001, //Do not include mxrt for NNEvo to auto-define a mutation rate
    }
    agents := network.NewNNEvo(&config) //initialize with params
    agents.CreatePopulation(nn)         //create network pool
    agents.Summary() //Print Agents Configuration

    inputs := [][]float64{
      {2.0, 1.0, 3.0, 4.0, 5.0},
      {-3.7, 4.7, -5.7, -6.7, 7.7},
      {14, 0.123, 0.89, -6.7, 7.0},
    }
    targets := [][]float64{
      {0, 1},
      {1, 0},
      {1, 0},
    }

    //Fit Data, Return best model
    //@Params: inputs, targets, validation inputs, validation outputs, loss function, verbosity
    model := agents.Fit(inputs, targets, nil, nil, "cross-entropy", 100)

    //Save Model
    filename := "example.model"
    model.Save(filename)
  }
```  

## Unsupervised Learning  
NNEvo uses an environment struct similar to OpenAI Gym environments. These are 
located in env.go. Once a environment is configured that implements DiscreteEnv 
or ContEnv, one can use NNEvo for unsupervised learning:  
```golang
  package main

  import (
    network "github.com/NN..." //import state for package 
    //other imports
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
    nn.Compile(false)
    nn.Summary()

    //create Env
    env := InitCustomEnv()

    //Configure NNEvo Params
    config := network.Config{
      Population:  100,
      Generations: 5000,
      Elites:      20,
      Goal:        100,
      Metric:      "reward",
      Callbacks:   []string{"checkpoint"},
    }
    agents := network.NewNNEvo(&config) //initialize with params
    agents.CreatePopulation(nn)         //create network pool
    agents.NewContEnv(env)              //Needed for agents to see environment
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
```  

# Examples  
| Dataset | Topology       | Generations | Accuracy | Link                                                     |
|---------|----------------|-------------|----------|----------------------------------------------------------|
| wines   | [13 16 16 8 3] | 10000       | 100%     | https://github.com/BlakeERichey/GoLang/tree/master/wines |
|         |                |             |          |                                                          |
|         |                |             |          |                                                          |