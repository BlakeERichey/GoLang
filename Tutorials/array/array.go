package main

import "fmt"

func main(){
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
	d = append(d, 3)	//adds 3 to end
	fmt.Println("d", d)
	fmt.Println(len(d))	//length
}