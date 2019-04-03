package brichey

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
