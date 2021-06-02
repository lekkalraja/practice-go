package main

import "fmt"

func main() {

	index := 0

	for index < 5 {
		fmt.Printf("Index Value : %d \n", index)
		index += 1
	}
}

/*
	While loops
		How does a while loop differ from a for loop?
		Ans) A for loop has a predefined amount of iterations. A while loop doesn’t necessarily.
	RUN:
		raja@raja-Latitude-3460:~/Documents/coding/golang/practice-go$ go run beginner-exercises/while-loops.go
		Index Value : 0
		Index Value : 1
		Index Value : 2
		Index Value : 3
		Index Value : 4
*/
