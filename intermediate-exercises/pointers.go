package main

import "fmt"

func main() {
	var name string = "Raja"
	var namePointer *string = &name
	fmt.Printf("Address location of name : %p == %p  where as pointer address is : %p\n",
		namePointer, &name, &namePointer)

	fmt.Printf("Value of name : %s == %s \n", *namePointer, name)
}

/*
	Pointers
	========
	1. Where are variables stored in the computer?
		* Variables will store on the RAM (Memory).
	2. What is a pointer?
		* Pointer is an variable, it can be used to point another variable by storing it's address location
			i.e. A pointer is a variable whose address is the direct memory address of another variable.
	3. How can you declare a pointer?
		* Address location of name : 0xc000010240 == 0xc000010240  where as pointer address is : 0xc00000e028
		* Value of name : Raja == Raja
*/
