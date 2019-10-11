package main

import (
	"fmt"
	"math/rand"
)

func main() {
	
	fmt.Println("Hello, playground")
	s:=createRndNums(1000, 2000, 1)
	fmt.Println("Original Array", s)
	fmt.Println("Comparisons:", selectionSort(s))
	fmt.Println("Array After:", s)
	
	s=createRndNums(1000, 2000, 2)
	fmt.Println("Original Array", s)
	fmt.Println("Comparisons:", bubbleSort(s))
	fmt.Println("Array After:", s)
}

func selectionSort(arr []int) int {
	counter := 0 //index of value to be swapped
	current := 0 //index of pointer searching unsorted array
	minIndex := 0 //min value's index while traversing unsorted array
	comparisons := 0 //number of comparisons in algorithm
	for counter < len(arr) {
		curVal:=arr[counter]
		for current < len(arr) {
			comparisons++
			if arr[current] < curVal && arr[current] < arr[minIndex] {
				minIndex = current
			}
			current++
		}
		
		//swap numbers
		swapping := arr[counter] //temp variable
		arr[counter]=arr[minIndex]
		arr[minIndex]=swapping
		counter++
		minIndex, current = counter, counter
		fmt.Println(counter, " Traversal ", "Min Index: ", minIndex, arr)
	}
	fmt.Println(arr)
	return comparisons
}

func bubbleSort(arr []int) int {
	comparisons := 0
	n := len(arr)
	
	updated := false
	for i:=0; i<n; i++ {
		for j:=0; j<n-i-1; j++{
			comparisons++
			if(arr[j] > arr[j+1]){
				arr[j+1], arr[j] = arr[j], arr[j+1]
				updated = true
			}
		}
		if(!updated){
			break
		}
	}
	fmt.Println(arr)
	return comparisons
}

//makes a list of size `size` and returns the generated list
func createRndNums(size int, maxNum int, seed int64) []int {
	rand.Seed(seed)
	s := make([]int, size)
	for i:=0;i < size;i++ {
		s[i] = rand.Intn(maxNum)
	}
	return s
}