package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFromConsole()
}

func readFromConsole() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")

	for {
		text := getText(reader)
		if strings.Compare(text, "Exit") == 0 {
			break
		}
		number := convertToInt(text)
		isBetween1And10(number)
	}
}

func getText(reader *bufio.Reader) string {
	fmt.Print("-> ")
	text, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal("Something went wrong", err)
	}

	text = strings.Replace(text, "\n", "", -1)
	return text
}

func convertToInt(text string) int {
	number, err := strconv.Atoi(text)
	if err == nil {
		fmt.Println(number)
	}
	return number

	/*i2, err := strconv.ParseInt(text, 10, 64)
	if err == nil {
		fmt.Println(i2)
	}*/

}

func isBetween1And10(number int) {
	if number >= 1 && number <= 10 {
		fmt.Printf("The Entered number is: %d  between 1 and 10 \n", number)
	} else {
		fmt.Printf("The Entered number is: %d not between 1 and 10 \n", number)
	}
}

/*
Keyboard input
	Get a number from the console and check if it’s between 1 and 10.

To Build & Execute:
===================
> raja@raja-Latitude-3460:~/Documents/coding/golang/go-exercises$ go build beginner-exercises/read-from-console.go
raja@raja-Latitude-3460:~/Documents/coding/golang/go-exercises$ ./read-from-console
Simple Shell
---------------------
-> 0
0
The Entered number is: 0 not between 1 and 10
-> 5
5
The Entered number is: 5  between 1 and 10
-> 10
10
The Entered number is: 10  between 1 and 10
-> 11
11
The Entered number is: 11 not between 1 and 10
-> Exit
raja@raja-Latitude-3460:~/Documents/coding/golang/go-exercises$

(or)

> raja@raja-Latitude-3460:~/Documents/coding/golang/go-exercises$ go run beginner-exercises/read-from-console.go
Simple Shell
---------------------
-> 5
5
The Entered number is: 5  between 1 and 10
-> 10
10
The Entered number is: 10  between 1 and 10
-> 11
11
The Entered number is: 11 not between 1 and 10
-> 0
0
The Entered number is: 0 not between 1 and 10
-> Exit
raja@raja-Latitude-3460:~/Documents/coding/golang/go-exercises$
*/

// Source : https://tutorialedge.net/golang/reading-console-input-golang/
