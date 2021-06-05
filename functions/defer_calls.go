package main

import "fmt"

func main() {
	testDefer()
}

func testDefer() {
	defer fmt.Println("Hello")
	defer fmt.Println("!")
	fmt.Println("World")
}

/*
	Defer
	=====
	Predict what this code does:
		defer fmt.Println("Hello")
		defer fmt.Println("!")
		fmt.Println("World")

	OUTPUT
	======
	raja@raja-Latitude-3460:~/Documents/coding/golang/practice-go$ go run functions/defer_calls.go
	World
	!
	Hello
	raja@raja-Latitude-3460:~/Documents/coding/golang/practice-go$ go run functions/defer_calls.go
	World
	!
	Hello
	raja@raja-Latitude-3460:~/Documents/coding/golang/practice-go$ go run functions/defer_calls.go
	World
	!
	Hello

	Observation: Printing defer call's in last in first out manner i.e. following Stack
*/
