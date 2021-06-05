package main

import "fmt"

func main() {
	name, age := getPerson()
	fmt.Printf("Person name : %s and age : %d \n", name, age)
}

func getPerson() (string, int) {
	return "Raja", 24
}

func getNumbers() (int, int) {
	return 23, 23
	//return "Raja", "Lekkala" // ERROR : cannot convert "Raja" (untyped string constant) to int
}

// In this case any return values can work, because we are not having specific return types
func getValues() (interface{}, interface{}) {
	//return 23, 23
	return "Raja", "Lekkala" // ERROR : cannot convert "Raja" (untyped string constant) to int
}

/*
	Multiple return
	================
		Change the return values from 2,4 to “hello”,”world”. Does it still work?
		 Ans) No, Go is statically typed language, so changing form 2,4 to "hello", "world" is
		 		same as changing data type from int to string, so in this case needs to change
				func return types from int, int to string, string

		Can a combination of strings and numbers be used?
		 Ans) Yes
*/
