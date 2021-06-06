package main

import (
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go printGreet("Raja", &wg)

	wg.Wait()
	log.Println("Done with the greeting!")
}

func printGreet(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Printf("Hi, %s \n", name)
}

/*
	Goroutines
	==========
	What is a goroutine?
		Ans) goroutine is a lightweight thread to execute some piece of independent code.
		     i.e. A goroutine is a function that can run concurrently.

	How can you turn a function into a goroutine?
		Ans) call function with `go` : `go func1()`

	What is concurrency?
		Ans) Concurrency, do lots of things at once.
		That’s different from doing lots of things at the same time
*/

/*
	OUTPUT:
	======
	raja@raja-Latitude-3460:~/Documents/coding/golang/practice-go$ go run goroutines_channels/goroutines.go
	2021/06/06 09:56:23 Hi, Raja
	2021/06/06 09:56:23 Done with the greeting!
*/
