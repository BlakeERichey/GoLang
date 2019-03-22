package main

import (
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//----------Reads all of file----------
	data, err := ioutil.ReadFile("test.txt")
	check(err)
	fmt.Println(string(data))

	//----------Read specific parts of file----------
	//f, err := os.Open("test.txt")
	//check(err)
	//...
}
