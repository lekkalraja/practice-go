package main

import (
	"fmt"
	"strings"
)

func main() {
	var greet string = "hello world"
	var greetSlice []string = strings.Split(greet, " ")
	fmt.Printf("Slice of Greet : %v \n", greetSlice)

	var countries [][]string = [][]string{
		{"India", "New Delhi"},
		{"Singapore", "Singapore"},
		{"Australia", "Sydney"},
	}

	fmt.Printf("Slice of Slices : %v \n", countries)
}

/*
	Slices
	======
	* Take the string ‘hello world’ and slice it in two.
	* Can you take a slice of a slice?
	RUN:
	===
	raja@raja-Latitude-3460:~/Documents/coding/golang/practice-go$ go run intermediate-exercises/slices.go
	Slice of Greet : [hello world]
	Slice of Slices : [[India New Delhi] [Singapore Singapore] [Australia Sydney]]
*/
