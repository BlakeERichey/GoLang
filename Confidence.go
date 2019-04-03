//Confidence Calculator for association rules

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	purchase := readString("Purchase: ")
	msg := "Support of " + purchase + ": "
	support := readNum(msg)
	buy := readString("Buy: ")
	msg = "Support of " + purchase + " and " + buy + ": "
	totalSup := readNum(msg)
	confidence := totalSup / support
	fmt.Println("Confidence: ", confidence)

}

func readNum(msg string) float64 {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(msg)
	scanner.Scan()
	val, _ := strconv.ParseFloat(scanner.Text(), 64)
	return val
}

func readString(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(msg)
	scanner.Scan()
	return scanner.Text()
}
