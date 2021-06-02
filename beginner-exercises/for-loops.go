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

To RUN:
=======
raja@raja-Latitude-3460:~/Documents/coding/golang/practice-go$ go run beginner-exercises/loops.go
 10 20 30 40
 20 40 60 80
 30 60 90 120
 40 80 120 160
iteration x: 1
iteration x: 2
iteration x: 3
iteration x: 4
iteration x: 5
iteration x: 6
iteration x: 7
iteration x: 8
iteration x: 9
iteration x: 10
Count : 1
Count : 2
Count : 3
Count : 4
Count : 5
Count : 6
Count : 7
Count : 8
Count : 9
Count : 10
raja@raja-Latitude-3460:~/Documents/coding/golang/practice-go$
*/
