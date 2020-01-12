package main

import (
	network "GoLang/NN"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

//implements ContEnv:

//ContEnv is used for NNEvo.Train and provides a replicatable methodology for evalauting
//a Networks fitness with a continuous action space
// type ContEnv interface {
// 	Step(action [][]float64) ([][]float64, float64, bool) //returns ob, reward, done
// 	Reset() [][]float64                                   //returns ob
// 	GetSteps() int
// 	Render()
// }

//StockEnv is an implementation of ContEnv that evaluates a NNs performance at
//guessing when to buy stocks
type StockEnv struct {
	steps                     int //how many steps have been taken
	done                      bool
	validate                  bool //return validation values
	inputs, targets           [][]float64
	validInputs, validTargets [][]float64
}

//InitStockEnv initializes a StockEnv with initial funds
func InitStockEnv() *StockEnv {
	env := new(StockEnv)
	env.steps = 0
	env.done = false
	env.validate = false
	inputs, targets := load("csvs")
	env.inputs, env.targets, env.validInputs, env.validTargets = network.ValidationSplit(inputs, targets, .2)
	return env
}

//Step takes and actions list corresponding to an action per observation
//performs corresponding actions from actionsList then returns
//caluclated results.
//returns (empty array, total rewards, done) empty array to comply with ContEnv
func (env *StockEnv) Step(actionsList [][]float64) ([][]float64, float64, bool) {
	var state [][]float64
	rewards := make([]float64, 0)
	for ob, arr := range actionsList {
		val := arr[0]
		var action string
		if val > 0.05 {
			action = "buy"
		} else if val < -0.05 {
			action = "sell"
		} else {
			action = "hold"
		}

		env.steps++
		if env.validate {
			env.done = env.steps >= len(env.inputs)
		} else {
			env.done = env.steps >= len(env.validInputs)
		}
		reward := env.getReward(ob, action)
		rewards = append(rewards, reward)
		// fmt.Println("Action:", action, "Reward:", reward)
		if env.done { //getReward terminated env
			break
		}
	}

	return state, sumArr(rewards...), env.done
}

//Render added to implement ContEnv. Shows relevant StocksEnv info
func (env *StockEnv) Render() {
	//Render info here
}

//Reset resets env.done and env.steps. Returns all observations
func (env *StockEnv) Reset() [][]float64 {
	env.done = false
	env.steps = 0
	env.validate = !env.validate
	if env.validate {
		return env.inputs
	}
	return env.validInputs
}

//GetSteps return number of steps taken in environment
func (env *StockEnv) GetSteps() int {
	return env.steps
}

func (env *StockEnv) getReward(ob int, action string) (reward float64) {
	answer := network.Argmax(env.targets[ob]...)
	if answer == 1 && action != "sell" {
		env.done = true
	} else if answer == 0 && action == "buy" {
		reward = 1.0
	}
	return
}

func sumArr(array ...float64) (total float64) {
	for _, val := range array {
		total += val
	}
	return total
}

func load(dir string) ([][]float64, [][]float64) {
	var files []string
	unableToLoadFiles := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if dir != path {
			files = append(files, path)
		}
		return nil
	})
	if unableToLoadFiles != nil {
		panic(unableToLoadFiles)
	}

	if len(files) < 50 {
		for _, file := range files {
			fmt.Println("Loaded", file)
		}
	}

	cummCsv := make([][]string, 0)
	for _, filename := range files {
		data, _ := ioutil.ReadFile(filename)
		rec := csv.NewReader(strings.NewReader(string(data)))
		records, err := rec.ReadAll()
		if err != nil {
			log.Fatal(err)
		}
		cummCsv = append(cummCsv, records[1:]...)
	}

	ids := make([]float64, len(cummCsv))
	inputs := make([][]float64, 0)
	// outputs := make([][]float64, 0)
	for i := range cummCsv {
		//row values as floats
		lstm := make([]float64, 50)
		data := make([]float64, 50)

		//convert values from string
		lstmarr := strings.Split(cummCsv[i][1][1:len(cummCsv[i][1])-1], ", ")
		dataarr := strings.Split(cummCsv[i][2][1:len(cummCsv[i][2])-1], ", ")
		for j, val := range lstmarr {
			//get lstm value
			datapoint, err := strconv.ParseFloat(val, 64)
			if err != nil {
				log.Fatal(err)
			}
			lstm[j] = datapoint

			//get data value
			datapoint, err = strconv.ParseFloat(dataarr[j], 64)
			if err != nil {
				log.Fatal(err)
			}
			data[j] = datapoint
		}
		//get closing value
		closing, err := strconv.ParseFloat(cummCsv[i][3], 64)
		if err != nil {
			log.Fatal(err)
		}
		if data[len(data)-1] < closing {
			ids[i] = 0
		} else {
			ids[i] = 1
		}
		inputs = append(inputs, lstm)
	}
	targets := onehot(2, ids)

	return inputs, targets
}

func onehot(classes int, data []float64) [][]float64 {
	encoded := make([][]float64, 0)
	for _, val := range data {
		categorical := make([]float64, classes)
		categorical[int(val)] = 1.0
		encoded = append(encoded, categorical)
	}
	return encoded
}
