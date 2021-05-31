package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Printf("Path of the Program : %s \n", args[:1])
	fmt.Printf("Program Arguments : %s \n", args[1:])
}

/*
Keyboard input
	Make a program that lets the user input a name

To Build & Execute:
===================
> go build beginner-exercises/command-line-args.go
> ./command-line-args Raja 23
Path of the Program : [./command-line-args]
Program Arguments : [Raja 23]
(or)
> go run beginner-exercises/command-line-args.go Raja 23
Path of the Program : [/tmp/go-build1405593461/b001/exe/command-line-args]
Program Arguments : [Raja 23]
*/
