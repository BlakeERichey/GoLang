package main

import "fmt"

func main() {
	first, last := "Blake", "Richey"
	fmt.Println(first == last) //false
	last = "Blake"
	fmt.Println(first == last) //true

	//indexing
	last = last[:len(last)-1]
	fmt.Println(last)                          //Blak
	fmt.Println(last[2])                       //97
	determineType, _ := fmt.Println(last[2:3]) //a
	fmt.Printf("Type is %T", determineType)    //int
	fmt.Println("")

	//substring
	fmt.Println(last[0:len(last)]) //Blak
	fmt.Println(last[0:4])         //Blak

}
