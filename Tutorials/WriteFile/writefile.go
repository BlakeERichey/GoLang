package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//with all contents
	d1 := []byte("Hello\nWorld\n")
	err := ioutil.WriteFile("./test.txt", d1, 0644)
	check(err)

	f, err := os.OpenFile("newtest.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600) //append to file
	check(err)

	defer f.Close()

	w := bufio.NewWriter(f)
	n3, err := w.WriteString("Adding this to file\n")
	fmt.Println("wrote %d bytes", n3)

	f.Sync()  //write to stable storage
	w.Flush() //ensure buffered operations have been applied
}
