package main

import "fmt"

func main() {
	weights := []int{20, 50, 65, 23, 87}

	totalWeight := 0
	for _, weight := range weights {
		totalWeight += weight
	}
	fmt.Printf("Average of (%v) : %d\n", weights, totalWeight/len(weights))
}

/*
Variables
Create a program that calculates the average weight of 5 people.

To Build & Execute:
===================
raja@raja-Latitude-3460:~/Documents/coding/golang/practice-go$ go build beginner-exercises/variables_1.go
raja@raja-Latitude-3460:~/Documents/coding/golang/practice-go$ ./variables_1
Average of ([20 50 65 23 87]) : 49
raja@raja-Latitude-3460:~/Documents/coding/golang/practice-go$ go run beginner-exercises/variables_1.go
Average of ([20 50 65 23 87]) : 49
raja@raja-Latitude-3460:~/Documents/coding/golang/practice-go$
*/
