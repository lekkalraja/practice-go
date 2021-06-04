package main

import "fmt"

type house struct {
	noRooms int
	price   float64
	city    string
}

func main() {
	mine := house{
		noRooms: 4,
		price:   10000000.00,
		city:    "Mysore",
	}

	fmt.Printf("My House : %+v \n", mine)
}

/*
	Struct
	=======
		Create a struct house with variables noRooms, price and city

		How does a struct differ from a class?
			* Classes are a concept of object oriented programming.
			  A class can be used to create objects. Classes can inherit from another.
			* Class will have both state and behaviors(encapsulation)
			  whereas struct will have state but needs to extend or associate behaviors
*/
