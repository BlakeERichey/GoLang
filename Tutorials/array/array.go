package main

import "fmt"

func main() {
	// var array [2]string
	// array[0] = "Hello"
	// array[1] = "World"
	// fmt.Println(array[0], array[1])	//hello world
	// fmt.Println(array)	//[hello world]

	// primes := [6]int{2,3,5,7,11,13}
	// fmt.Println(primes)

	// //slicing does not copy
	// var s []int = primes[0:3] //2,3,5
	// fmt.Println(s)

	// s[0] = 0
	// fmt.Println(primes) //0,3,5,7,11,13

	//-----make slice feature-----
	//The make function allocates a zeroed array and returns a slice that
	//refers to that array:
	//allows for dynamic arrays
	a := make([]int, 5)
	fmt.Println("a", a)

	b := make([]int, 0, 5)
	fmt.Println("b", b)

	c := b[:2]
	fmt.Println("c", c)

	d := c[2:5]
	d = append(d, 3) //adds 3 to end
	fmt.Println("d", d)
	fmt.Println(len(d)) //length

	// Creating and initializing an array
	// Using shorthand declaration
	my_arr1 := [6]int{12, 456, 67, 65, 34, 34}

	// Copying the array into new variable
	// Here, the elements are passed by reference
	my_arr2 := &my_arr1

	fmt.Println("Array_1: ", my_arr1)
	fmt.Println("Array_2:", *my_arr2)

	my_arr2[5] = 1000000

	// Here, when we copy an array
	// into another array by reference
	// then the changes made in original
	// array will reflect in the
	// copy of that array
	fmt.Println("\nArray_1: ", my_arr1)
	fmt.Println("Array_2:", *my_arr2)

	resetArray((*my_arr2)[:]...)
	fmt.Println("\nArray_1: ", my_arr1)
	fmt.Println("Array_2:", *my_arr2)
	fmt.Println("Array_2:", len(*my_arr2))
}

func resetArray(test ...int) {
	for i := range test {
		test[i] = 0
	}
}
