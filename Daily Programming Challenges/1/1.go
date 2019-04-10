package main

import "fmt"

func main() {

	list := []int{3, 17, 5, 9, 7, 6, 8}
	fmt.Println(sumToVal(list, 18))

}

func sumToVal(list []int, sum int) bool {
	for index, val := range list {
		if contains(list[index+1:], sum-val) {
			return true
		}
	}
	return false
}

func contains(list []int, val int) bool {
	for _, ele := range list {
		if ele == val {
			return true
		}
	}
	return false
}
