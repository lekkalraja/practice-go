package main

import "fmt"

func main() {
	printSum(2.4, 3)
	printSum(4, 3.6)
}

func printSum(num1 float64, num2 float64) {
	sum := add(num1, num2)
	fmt.Printf("Sum of %f & %f : %f \n", num1, num2, sum)
	fmt.Printf("Sum of %.2f & %.2f : %.2f \n", num1, num2, sum)
}

func add(num1, num2 float64) float64 {
	return num1 + num2
}

/*
	Methods
		Create a method that sums two numbers
		Create a method that calls another method.

	OUTPUT:
		raja@raja-Latitude-3460:~/Documents/coding/golang/practice-go$ go run functions/methods.go
		Sum of 2.400000 & 3.000000 : 5.400000
		Sum of 2.40 & 3.00 : 5.40

		Sum of 4.000000 & 3.600000 : 7.600000
		Sum of 4.00 & 3.60 : 7.60
*/
