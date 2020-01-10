package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//get stats
	stats := getStats()
	//fmt.Println(stats)

	//get best stats
	count := 0
	for avg(stats) < 17.1 {
		count++
		stats = getStats()
	}
	fmt.Println(stats, avg(stats), count) //[17, 18, 15, 18, 17, 18]
}

func roll() int {
	return rand.Intn(6) + 1
}

func getAttribute() (rolls []int) {
	rolls = make([]int, 4)
	for i := 0; i < 4; i++ {
		rolls[i] = roll()
	}

	//find min
	min := rolls[0]
	index := 0
	for i, v := range rolls {
		if v < min {
			min = v
			index = i
		}
	}

	//remove min
	rolls[index] = rolls[len(rolls)-1]
	rolls = rolls[:len(rolls)-1]

	return
}

func getStats() (stats []int){
	stats = make([]int, 6)
	for i, _ := range stats {
		stats[i] = sum(getAttribute())
	}
	return
}

func sum(arr []int) (total int) {
	total = 0
	for _, v := range arr {
		total += v
	}
	return
}

func avg(arr []int) float64 {
	return float64(sum(arr)) / float64(len(arr))
}

