package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	number, err := strconv.Atoi(os.Args[1])

	if err != nil {
		log.Fatalf("Failed to Process : %v", os.Args[1])
	}

	if number > 0 {
		if number%2 == 0 {
			fmt.Printf("Number : %d is an even number \n", number)
		} else {
			fmt.Printf("Number : %d is an even number \n", number)
		}
	} else {
		fmt.Printf("Number is %d, so not processing \n", number)
	}
}

/*
	If statements
		Make a program that divides x by 2 if it’s greater than 0
		Find out if if-statements can be used inside if-statements.

	To RUN :
	=======
	raja@raja-Latitude-3460:~/Documents/coding/golang/practice-go$ go run beginner-exercises/if-statements.go 0
	Number is 0, so not processing
	raja@raja-Latitude-3460:~/Documents/coding/golang/practice-go$ go run beginner-exercises/if-statements.go 4
	Number : 4 is an even number
	raja@raja-Latitude-3460:~/Documents/coding/golang/practice-go$ go run beginner-exercises/if-statements.go 7
	Number : 7 is an even number
	raja@raja-Latitude-3460:~/Documents/coding/golang/practice-go$ go run beginner-exercises/if-statements.go ad
	2021/06/02 08:54:10 Failed to Process : ad
	exit status 1
*/
