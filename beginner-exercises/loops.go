package main

import "fmt"

func main() {

	// Loop inside loop
	for _, num := range []int{1, 2, 3, 4} {
		for _, dig := range []int{10, 20, 30, 40} {
			fmt.Printf(" %d", num*dig)
		}
		fmt.Println()
	}

	for x := 1; x <= 10; x++ {
		fmt.Printf("iteration x: %d\n", x)
	}

	for index := range [10]int{} {
		fmt.Printf("Count : %d \n", index+1)
	}
}

/*
	For loops
		Can for loops exist inside for loops? -> Yes
		Make a program that counts from 1 to 10.
*/
