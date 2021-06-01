package main

import "fmt"

var global string = "I am global"

func main() {
	var local string = "I am local"
	fmt.Printf("Global Variable in main : %s \n", global)
	fmt.Printf("Local Variable in main : %s \n", local)
	helper()
}

func helper() {
	fmt.Printf("Global Variable in helper : %s \n", global)
	// fmt.Printf("Local Variable in main : %s \n", local) // Can't find local variable
}

/*
 Scope
	* What’s the difference between a local and global variable?
		* Local variable will only visible to that func or block or loop or section
		whereas global variable can be accessed anywhere in the package where it is defined.
	How can you make a global variable?
		* Global variable will be define outside of any block or function.

To RUN:
=======
raja@raja-Latitude-3460:~/Documents/coding/golang/practice-go$ go run beginner-exercises/scope.go
Global Variable in main : I am global
Local Variable in main : I am local
Global Variable in helper : I am global
raja@raja-Latitude-3460:~/Documents/coding/golang/practice-go$
*/
