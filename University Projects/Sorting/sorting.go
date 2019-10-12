//Blake Richey
//COSC 3325 - Dr Rainwater
//Implement 3 sorting algorithms and empirically evaluate their performance
//This code generates a randomly generated list and tests each algorithm
//on different sizeds lists. Reports the number of comparisons in each algorithm
//and displays them to the user

package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	var seed int64 //set seed to 0
	maxNum := 50000

	//-------------------- Selection Sort --------------------
	fmt.Println("SelectionSort Results:")
	s := createRndNums(10, maxNum, seed)
	fmt.Println("Number of values in array", 10)
	fmt.Println("Number of comparisons required:", selectionSort(s))

	s = createRndNums(100, maxNum, seed)
	fmt.Println("Number of values in array", 100)
	fmt.Println("Number of comparisons required:", selectionSort(s))

	s = createRndNums(1000, maxNum, seed)
	fmt.Println("Number of values in array", 1000)
	fmt.Println("Number of comparisons required:", selectionSort(s))

	s = createRndNums(10000, maxNum, seed)
	fmt.Println("Number of values in array", 10000)
	fmt.Println("Number of comparisons required:", selectionSort(s))

	//-------------------- Bubble Sort --------------------
	fmt.Println("\nBubbleSort Results:")
	s = createRndNums(10, maxNum, seed)
	fmt.Println("Number of values in array", 10)
	fmt.Println("Number of comparisons required:", bubbleSort(s))

	s = createRndNums(100, maxNum, seed)
	fmt.Println("Number of values in array", 100)
	fmt.Println("Number of comparisons required:", bubbleSort(s))

	s = createRndNums(1000, maxNum, seed)
	fmt.Println("Number of values in array", 1000)
	fmt.Println("Number of comparisons required:", bubbleSort(s))

	s = createRndNums(10000, maxNum, seed)
	fmt.Println("Number of values in array", 10000)
	fmt.Println("Number of comparisons required:", bubbleSort(s))

	//-------------------- Comb Sort --------------------
	fmt.Println("\nCombSort Results:")
	s = createRndNums(10, maxNum, seed)
	fmt.Println("Number of values in array", 10)
	fmt.Println("Number of comparisons required:", combSort(s))

	s = createRndNums(100, maxNum, seed)
	fmt.Println("Number of values in array", 100)
	fmt.Println("Number of comparisons required:", combSort(s))

	s = createRndNums(1000, maxNum, seed)
	fmt.Println("Number of values in array", 1000)
	fmt.Println("Number of comparisons required:", combSort(s))

	s = createRndNums(10000, maxNum, seed)
	fmt.Println("Number of values in array", 10000)
	fmt.Println("Number of comparisons required:", combSort(s))

	fmt.Println("\nHit Enter to end the program...")
	var input string
	fmt.Scanln(&input) //hit enter to end program
}

func selectionSort(arr []int) int {
	n := len(arr)

	//current: currently observed index
	//counter: arr is correctly sorted to this index
	//comparison: how many if comparisons have been made
	current, counter, comparisons := 0, 0, 0
	for counter = 0; counter < n; counter++ {

		//search for smallest unsorted number
		minIndex := counter
		for current = counter + 1; current < n; current++ {
			comparisons++
			if arr[current] < arr[minIndex] {
				minIndex = current
			}
		}

		//put minVal into counters location
		arr[counter], arr[minIndex] = arr[minIndex], arr[counter]
	}
	return comparisons
}

func bubbleSort(arr []int) int {
	comparisons := 0
	n := len(arr)

	updated := false
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			comparisons++
			if arr[j] > arr[j+1] { //if arr[j] > arr[j+1] swap the two values
				arr[j+1], arr[j] = arr[j], arr[j+1]
				updated = true
			}
		}

		//break if no changes were implemented, aka: arr is sorted
		if !updated {
			break
		}
		updated = false
	}
	return comparisons
}

func combSort(arr []int) int {
	gap, swapped, comparisons := len(arr), true, 0
	for swapped || gap > 1 { //for = golang's while loop
		gap = int(math.Max(float64(gap)/1.3, 1))
		swapped = false
		for i := 0; i < len(arr)-gap; i++ {
			j := i + gap
			comparisons++
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i] //swap values
				swapped = true
			}
		}
	}

	return comparisons
}

//makes a list of size `size` and returns the generated list
func createRndNums(size int, maxNum int, seed int64) []int {
	rand.Seed(seed)
	s := make([]int, size)
	for i := 0; i < size; i++ {
		s[i] = rand.Intn(maxNum)
	}
	return s
}
