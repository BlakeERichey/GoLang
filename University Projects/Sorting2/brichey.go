//Blake Richey
//COSC 3325 - Dr Rainwater
//Implement 3 sorting algorithms and empirically evaluate their performance
//This code generates a randomly generated list and tests each algorithm
//on different sizeds lists. Reports the number of comparisons in each algorithm
//and displays them to the user

package main

import (
	"fmt"
	//"math"
	"math/rand"
)

var msComps, qsComps int

func main() {
	fmt.Println("Sorting Algorithm: Insertion Sort")
	makeComps("insertion", 10)
	fmt.Println("\nSorting Algorithm: Merge Sort")
	makeComps("merge", 10)
	fmt.Println("\nSorting Algorithm: Quick Sort")
	makeComps("quick", 10)

}

func makeComps(ftype string, numVals int) string {
	if numVals > 10000 {
		return ""
	}

	unsorted := createRndNums(numVals)
	if ftype == "insertion" {
		fmt.Println("Number of values in array:", numVals)
		fmt.Println("Number of comparisons:", insertionSort(unsorted))
	} else if ftype == "merge" {
		fmt.Println("Number of values in array:", numVals)
		msComps = 0
		mergesort(unsorted, 0, len(unsorted)-1)
		fmt.Println("Number of comparisons:", msComps)
	} else if ftype == "quick" {
		fmt.Println("Number of values in array:", numVals)
		qsComps = 0
		quicksort(unsorted)
		fmt.Println("Number of comparisons:", qsComps)
	}

	return makeComps(ftype, numVals*10)
}

func insertionSort(arr []int) int {
	comparisons := 0
	for i := 1; i < len(arr); i++ {
		val := arr[i]
		pointer := i - 1 //points to ele to compare aginst val to determine if swap is needed
		for pointer >= 0 {
			comparisons++
			temp := arr[pointer]
			if temp > val {
				arr[pointer+1], arr[pointer] = arr[pointer], arr[pointer+1] //swap val leftward
			}
			pointer--
		}
	}
	//fmt.Println(arr) //print sorted list
	return comparisons
}

//makes a list of size `size` and returns the generated list
func createRndNums(size int) []int {
	rand.Seed(0)
	maxNum := 50000
	s := make([]int, size)
	for i := 0; i < size; i++ {
		s[i] = rand.Intn(maxNum)
	}
	return s
}

func mergesort(array []int, lower int, upper int) {
	var middle int
	if lower < upper {
		middle = (lower + upper) / 2

		mergesort(array, lower, middle)
		mergesort(array, middle+1, upper)
		merge(array, lower, middle, upper)
	}
}

func merge(array []int, l, m, u int) {
	L := make([]int, m-l+1)
	R := make([]int, u-m)

	//fmt.Println(array, l,m,u)
	index := 0
	for index < len(L) {
		L[index] = array[l+index]
		index++
	}
	//fmt.Println("L", L)

	index = 0
	for index < len(R) {
		R[index] = array[m+1+index]
		index++
	}
	//fmt.Println("R", R)

	i, j, k := 0, 0, l

	for i < len(L) && j < len(R) {
		msComps++
		if L[i] < R[j] {
			array[k] = L[i]
			k++
			i++
		} else {
			array[k] = R[j]
			k++
			j++
		}
	}

	for i < len(L) {
		array[k] = L[i]
		i++
		k++
	}
	for j < len(R) {
		array[k] = R[j]
		k++
		j++
	}

}

func quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1

	pivot := rand.Int() % len(arr)

	arr[pivot], arr[right] = arr[right], arr[pivot]

	for i, _ := range arr {
		qsComps++
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}

	//reassign pivot to correct location
	arr[left], arr[right] = arr[right], arr[left]

	//sort subarrays
	quicksort(arr[:left])
	quicksort(arr[left+1:])

	return arr
}
