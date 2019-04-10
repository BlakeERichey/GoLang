package main

import (
	brichey "GoLang/Library"
	"fmt"
)

func main() {
	alpha := brichey.ReadNum("Enter Alpha: ")
	allNode := brichey.ReadNum("Enter total number of nodes: ")
	conNodes := brichey.ReadNum("Enter number of nodes connecting to node: ")

	fmt.Println("Probability: ", alpha/allNode+(1-alpha)*1/conNodes)
}
