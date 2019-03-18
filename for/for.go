package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//init and post is optional, this is equivalent to while
	sum := 1
	for sum < 10 {
		sum += sum
	}
	fmt.Println(sum)

	//can omit condition to loop forever
	//for {
	//	fmt.Println("printing...")
	//}
}
