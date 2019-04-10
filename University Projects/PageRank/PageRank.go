package main

import (
	"fmt"
)

func main() {
	lp := [6][6]float64{
		{0, 1, 0, 0, 0, 0},
		{0, 0, 0, .5, 0, .5},
		{1 / 3.0, 1 / 3.0, 0, 1 / 3.0, 0, 0},
		{0, 0, 0, 0, 0, 1},
		{1 / 4.0, 0, 1 / 4.0, 1 / 4.0, 0, 1 / 4.0},
		{0, 0, 0, 0, 0, 0},
	}
	fmt.Println(lp)
}

func pageRank(node int, alpha float64, list [][]float64) {
	cols := len(list[0]) //number of elements in row = cols
	rank := alpha / float64(cols)
	for i := 0; i < cols; i++ {
		rank += (1 - alpha) * 1 //need to figure out p(i)
	}
}
