package main

import "fmt"

func main() {
	print("Janaki", "Ram", "Radha", "Krishna")
	cities := []string{"Hyderabad", "Chennai", "Bangalore"}
	print(cities...)
}

func print(names ...string) {
	fmt.Println(names)
}

/*
	Variadic functions
	==================
		Create a variadic function that prints the names of students

	OUTPUT:
	raja@raja-Latitude-3460:~/Documents/coding/golang/practice-go$ go run functions/variadic_functions.go
	[Janaki Ram Radha Krishna]
	[Hyderabad Chennai Bangalore]
*/
