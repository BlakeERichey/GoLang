package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4, 5}
	fmt.Println(multipleArr(list))
}

//moves first index to last
func shift(list []int) []int {
	rv := list[1:]
	rv = append(rv, list[0])
	return rv
}

//returns product of all ele in list, ignores first ele
func totalMultiple(list []int) int {
	product := 1
	len := len(list)
	for index := 1; index < len; index++ {
		product *= list[index]
	}
	return product
}

func multipleArr(list []int) []int {
	rv := make([]int, len(list))
	for i := range list {
		rv[i] = totalMultiple(list)
		list = shift(list)
	}
	return rv
}
