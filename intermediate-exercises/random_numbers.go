package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	number := r.Intn(5) + 1

	fmt.Printf("Roll value : %d \n", number)

	// Another Way
	rand.Seed(time.Now().UnixNano())

	randomNum := random(1, 7)
	fmt.Printf("Random number: %d\n", randomNum)

	negativeNum := random(-7, 7)
	fmt.Printf("Negative Random number: %d\n", negativeNum)
}

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

/*
	Random numbers
	===============
	* Make a program that rolls a dice (1 to 6)
	* Can you generate negative numbers?
		* Yes

*/
