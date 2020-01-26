package main

import (
	"fmt"

	"github.com/red1bluelost/OL_Introduction/link"
)

var GLinker link.Linker

func main() {
	//create test linker pointer to the global variable
	linkTester := new(link.Linker)
	linkTester = &GLinker
	linkTester.Initialize()
	linkTester.Reset()

	//get test input
	var input string
	fmt.Scan(&input)

	//output result
	fmt.Print(input)
}
