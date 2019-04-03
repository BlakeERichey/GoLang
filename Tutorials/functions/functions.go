package main

import "fmt"

func main(){
	y := add(5,5)
	fmt.Println(y)
	fmt.Println(swap("World", "Hello"))
}

// func add(x int, y int) int{
// 	return x+y
// }

//if types are same can do like this:
func add(x, y int) int{
	return x+y
}

//multiple returned values
func swap(x string, y string) (string, string) {
	return y, x
}