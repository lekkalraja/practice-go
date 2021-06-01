package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	names := []string{"Janaki", "Ram", "Radha", "Krishna"}

	fmt.Printf("Numbers : %v \n", numbers)
	fmt.Printf("Names : %v \n", names)
}

/*
	Arrays
		Create an array with the number 0 to 10
		Create an array of strings with names

To Run:
=======
raja@raja-Latitude-3460:~/Documents/coding/golang/practice-go$ go run beginner-exercises/arrays.go
Numbers : [0 1 2 3 4 5 6 7 8 9 10]
Names : [Janaki Ram Radha Krishna]
raja@raja-Latitude-3460:~/Documents/coding/golang/practice-go$

*/
