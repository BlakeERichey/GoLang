package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//get stats
	stats := getStats()
	fmt.Println(stats)
	
	//get best stats
	count := 0
	for avg(stats) < 17{
		count++
		stats = getStats()
	}
	fmt.Println(stats, avg(stats), count)
}

func roll() int {
	return rand.Intn(6) + 1
}

func getAttribute() []int{
	rolls := make([]int, 4)
	for i:=0; i<4; i++ {
		rolls[i] = roll()
	}
	
	//find min
	min := rolls[0]
	index := 0
	for i, v := range rolls{
		if v < min{
			min = v
			index = i
		}
	}
	
	//remove min
	rolls[index] = rolls[len(rolls)-1]
	rolls = rolls[:len(rolls)-1]
	
	return rolls	
}

func getStats() []int {
	stats := make([]int, 6)
	for i, _ := range stats {
		stats[i] = sum(getAttribute())
	}
	return stats
}

func sum(arr []int) int {
	total := 0
	for _, v := range arr {
		total += v
	}
	return total
}

func avg(arr []int) float64 {
	return float64(sum(arr)) / float64(len(arr))
}
