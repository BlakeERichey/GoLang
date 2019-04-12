package brichey

//import statement: brichey "GoLang/Library"

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//reads a float from console and returns it
func ReadNum(msg string) float64 {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(msg)
	scanner.Scan()
	val, _ := strconv.ParseFloat(scanner.Text(), 64)
	return val
}

//reads a string from the console and returns it
func ReadString(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(msg)
	scanner.Scan()
	return scanner.Text()
}

//Contains: returns true is list contains an element == val
//list: list of values to look at
//val:	val to compare elements of list to
func Contains(list []int, val int) bool {
	for _, ele := range list {
		if ele == val {
			return true
		}
	}
	return false
}

//moves first index to last
func Shift(list []int) []int {
	rv := list[1:]
	rv = append(rv, list[0])
	return rv
}
