package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)


func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------")

	for{
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')

		text = strings.Replace(text, "\r\n", "", -1)
		fmt.Println(text)
		fmt.Println("hi" == text)
		if strings.Compare("hi", text) == 0 {
			fmt.Println("Hello, yourself")
		}
	}
}