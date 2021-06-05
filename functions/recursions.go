package main

import "fmt"

func main() {
	fmt.Printf("Sum till  %d : %d\n", 5, sumUp(5, 0))
	fmt.Printf("Sum till  %d : %d\n", 3, sumUp(3, 0))
}

func sumUp(num int, sum int) int {
	if num == 0 { // break point
		return sum
	}
	return sumUp(num-1, sum+num) // callig itself
}

/*
	Recursion
	When is a function recursive?
		Ans) when a function call's itself is called recursion
	Can a recursive function call non-recursive functions?
		Ans) Yes

	OUTPUT:
	======
	raja@raja-Latitude-3460:~/Documents/coding/golang/practice-go$ go run functions/recursions.go
	Sum till  5 : 15
	Sum till  3 : 6
*/
