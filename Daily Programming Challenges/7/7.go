package main

import (
	brichey "GoLang/Library"
	"fmt"
	"strconv"
)

func main() {
	mapping := make([]string, 26)
	for i := range mapping { //[a, b, c..., z]
		mapping[i] = string(i + 97)
	}

	fmt.Println(mapping)
	count := 0
	encodedString := "4467638"

	//-----First Pass-----
	valid := true
	validNums := make([]int, 9)
	for i := range validNums { //[1-9]
		validNums[i] = i + 1
	}

	//if substrings can be constructed out of validNums
	for i := range encodedString {
		num, _ := strconv.Atoi((encodedString[i : i+1]))
		if !brichey.Contains(validNums, num) {
			valid = false
		}
	}
	if valid {
		count++
	}

	fmt.Println("Count after First pass", count)

	//-----Second Pass-----
	valid = true
	if len(encodedString)%2 == 0 { //if string is even in length
		validNums = make([]int, 17)
		for i := range validNums { //[10-26]
			validNums[i] = i + 10
		}

		//if substrings can be constructed out of validNums
		for i := 0; i < len(encodedString)-1; i += 2 {
			num, _ := strconv.Atoi((encodedString[i : i+2]))
			if !brichey.Contains(validNums, num) {
				valid = false
			}
		}
	} else {
		valid = false
	}
	if valid {
		count++
	}

	fmt.Println("Count after Second pass", count)
}

//func passOne(string, list []string)
