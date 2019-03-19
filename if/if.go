package main

import "fmt"

func main(){
	//x:=0
	// if x < 3 {
	// 	fmt.Println("Hello")
	// }

	// for x < 12 {
	// 	if x == 3 { fmt.Println(x) }
	// 	x++
	// }

	// if y:=0; y>1 {
	// 	y++
	// 	fmt.Println(y)	//1
	// }else{
	// 	fmt.Println(y-1)	//-1
	// }

	//-----defer statements-----
	defer fmt.Println("World")
	fmt.Println("Hello")
	fmt.Println("From, Blake")
	//Prints hello from, blake world
	//defer happens after main ends
}