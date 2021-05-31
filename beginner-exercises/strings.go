package main

import "fmt"

var (
	gVar1 string // Default Value ""
	gVar2 int    // Default Value `0`
)

var gVar3 string = "GOLANG"
var gVar4 = "GO" // Compiller will not throw error even if we don't use declared global variable

// globalVariable5 := "GO-LLANG" // Invalid Declaration

func main() {

	var (
		lVar1 string = "Raja"
		//lVar2 := 23 // Invalid Declaration
		lVar2 int // Default Value 0
	)

	lVar3 := 23
	var lVar4 string = "Singapore"

	fmt.Println(lVar1, lVar2, lVar3, lVar4) // Raja 0 23 Singapore
	fmt.Println(gVar1, gVar2, gVar3)        // "" 0 GOLANG

}

/*
Strings
	Create a program with multiple string variables
	Create a program that holds your name in a string.

To Build & Execute:
===================
> go build beginner-exercises/strings.go
> ./strings
(or)
> go run beginner-exercises/strings.go
*/
