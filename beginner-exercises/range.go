package main

import "fmt"

func main() {

	array := [3]string{"one", "two", "three"}

	for index, element := range array {
		fmt.Printf("[ARRAY] %d -> %s \n", index, element)
	}

	// array[4] = "four" // index 4 (constant of type int) is out of bounds

	slice := []string{"one", "two", "three"}

	for index, element := range slice {
		fmt.Printf("[SLICE] %d -> %s \n", index, element)
	}

	/*slice[3] = "four"

	fmt.Printf("[SLICE] %d -> %s \n", 3, slice[3]) */
}

/*
	Range
	* What is the purpose of range ?
		* Range can be used to iterate over list of elements (slice, array, map, ... etc)
		* What the difference between the line
			1 ) for index, element := range a and -> Here we are initializing index and element from a
			2 ) the line for _, element := range a ? -> Here we are initializing only element and ignoring index.
*/
