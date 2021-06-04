package main

import "fmt"

func main() {
	var states map[string]string
	states = make(map[string]string)
	states["AP"] = "Andhra Pradesh"
	states["TN"] = "TamilNadu"

	fmt.Printf("States : %+v \n", states)

	cities := map[string]string{
		"HYD": "Hyderabad",
		"CHN": "Chennai",
		"BLR": "Bangalore",
	}
	fmt.Printf("Cities : %+v \n", cities)
}

/*
	Maps
	====
	* What is a map?
		* Map is a datastructure, it can be used to store data as key-value pairs
		  i.e. A Golang map is a unordered collection of key-value pairs.
	* Is a map ordered?
		* No, map will store data based on keys, so insertion/sorting order will not preserve
	* What can you use a map for?
		* Can be used to store data as key-value pairs
*/
