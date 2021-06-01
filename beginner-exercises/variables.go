package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	const shortForm = "2006-Jan-02"
	dob := os.Args[1]
	dobTime, _ := time.Parse(shortForm, dob)
	fmt.Printf("Year : %d  (%s)\n", dobTime.Year(), dobTime.String())
	fmt.Printf("Age : %d \n", time.Now().Year()-dobTime.Year())
}

/*
 Variables
	Calculate the year given the date of birth and age

To Build & Execute:
===================
raja@raja-Latitude-3460:~/Documents/coding/golang/practice-go$ go build beginner-exercises/variables.go
raja@raja-Latitude-3460:~/Documents/coding/golang/practice-go$ ./variables 1993-May-01
Year : 1993  (1993-05-01 00:00:00 +0000 UTC)
Age : 28
(or)
raja@raja-Latitude-3460:~/Documents/coding/golang/practice-go$ go run beginner-exercises/variables.go 1993-May-01
Year : 1993  (1993-05-01 00:00:00 +0000 UTC)
Age : 28
raja@raja-Latitude-3460:~/Documents/coding/golang/practice-go$
*/
