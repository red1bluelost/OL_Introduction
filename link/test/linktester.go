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
	linkTester.Reset()

	for {
		//get test input
		var input string
		fmt.Scan(&input)
		if input == "quit" {break}

		//handle link assignment
		linkTester.Handle(input)

		//output result
		linkTester.DebugLinks()
	}
}
