package network

type NNEvo struct {
	population  []Network
	generations int
	elites      int
	mxrt        float64
	goal        float64
	validation  bool
	env         Env
}

type Env interface {
	init()
	step() []float64
	reset() []float64
	getSteps() int
	getReward() float64
}

func NewNNevo() *NNEvo {
	return new(NNEvo)
}

//createPopulation creates duplicates of provided network
func (agents *NNEvo) createPopulation(nn *Network) {

}
