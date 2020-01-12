package network

import (
	"gonum.org/v1/gonum/mat"
)

//Worker is an advanced agent utilized by a genetic algorithm
type Worker struct {
	net       *Network
	mutations float64
	patience  int       //how long before discarding genes
	priority  int       //outlier in fitness trend?
	history   int       //history of priorities
	mask      mat.Dense //mask to apply to a layer
	//Environment for RL
	disc DiscreteEnv
	cont ContEnv
}

//Fitness()
//Mutate()
//Breed(b)
