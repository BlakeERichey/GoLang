package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	list := []int{3, 5, 1, 2, 6, 9, 8, 11, 10, 7, 4, 7, 7, 1, 2, 9}
	fmt.Println("First missing is", findMissingInt(list))
}

func findMissingInt(list []int) int {
	sort.Ints(list)
	for i, ele := range list[:len(list)-1] {
		if math.Abs(float64(list[i+1]-ele)) >= 2 {
			return ele + 1
		}
	}
	return list[len(list)-1] + 1
}
