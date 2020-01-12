package network

//DiscreteEnv is used for NNEvo.Train and provides a replicatable methodology for evalauting
//a Networks fitness with a discrete action space
type DiscreteEnv interface {
	Step(action int) ([][]float64, float64, bool) //returns ob, reward, done
	Reset() [][]float64                           //returns ob
	GetSteps() int
	Render()
}

//ContEnv is used for NNEvo.Train and provides a replicatable methodology for evalauting
//a Networks fitness with a continuous action space
type ContEnv interface {
	Step(action [][]float64) ([][]float64, float64, bool) //returns ob, reward, done
	Reset() [][]float64                                   //returns ob
	GetSteps() int
	Render()
}

//RunDiscrete runs a DiscreteEnv number of times equal to episodes and returns
//the average rewards for each episode.
//If validate == true, will run another episodes number of instance and
//return the validation results
//returns: rewards, validation rewards
func RunDiscrete(env DiscreteEnv, nn *Network, episodes int, validate, render bool) (float64, float64) {
	//initial run
	totalRewards := make([]float64, episodes)
	validateRewards := make([]float64, episodes)
	for i := 0; i < episodes; i++ {
		done := false
		envstate := env.Reset()
		rewards := make([]float64, 0)
		for !done {
			state, reward, end := env.Step(Argmax(nn.FeedForward(envstate)[0]...))
			envstate, done = state, end
			rewards = append(rewards, reward)
			if render {
				env.Render()
			}
		}
		totalRewards[i] = sumArr(rewards...)
	}

	//validation run
	if validate {
		for i := 0; i < episodes; i++ {
			done := false
			envstate := env.Reset()
			rewards := make([]float64, 0)
			for !done {
				state, reward, end := env.Step(Argmax(nn.FeedForward(envstate)[0]...))
				envstate, done = state, end
				rewards = append(rewards, reward)
				if render {
					env.Render()
				}
			}
			validateRewards[i] = sumArr(rewards...)
		}
	}

	if validate {
		n := float64(episodes)
		return sumArr(totalRewards...) / n, sumArr(validateRewards...) / n
	}
	return sumArr(totalRewards...) / float64(episodes), 0.0
}

//RunCont runs a ContEnv number of times equal to episodes and returns
//the average rewards for each episode.
//If validate == true, will run another episodes number of instance and
//return the validation results
//returns: rewards, validation rewards
func RunCont(env ContEnv, nn *Network, episodes int, validate, render bool) (float64, float64) {
	totalRewards := make([]float64, episodes)
	validateRewards := make([]float64, episodes)
	for i := 0; i < episodes; i++ {
		done := false
		envstate := env.Reset()
		rewards := make([]float64, 0)
		for !done {
			state, reward, end := env.Step(nn.FeedForward(envstate))
			envstate, done = state, end
			rewards = append(rewards, reward)
			if render {
				env.Render()
			}
		}
		totalRewards[i] = sumArr(rewards...)
	}

	if validate {
		for i := 0; i < episodes; i++ {
			done := false
			envstate := env.Reset()
			rewards := make([]float64, 0)
			for !done {
				state, reward, end := env.Step(nn.FeedForward(envstate))
				envstate, done = state, end
				rewards = append(rewards, reward)
				if render {
					env.Render()
				}
			}
			validateRewards[i] = sumArr(rewards...)
		}
	}

	if validate {
		n := float64(episodes)
		return sumArr(totalRewards...) / n, sumArr(validateRewards...) / n
	}
	return sumArr(totalRewards...) / float64(episodes), 0.0
}
