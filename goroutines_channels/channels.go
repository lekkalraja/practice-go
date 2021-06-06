package main

import "log"

func main() {
	ch := make(chan string)

	go reply("Raja", ch)

	log.Printf("[main-routine] Go-Routine communicated : %s\n", <-ch)
}

func reply(name string, ch chan<- string) {
	log.Println("[Go-routine] sending reply to main-routine")
	ch <- "Hello, " + name
}

/*
	Channels
	========
	When do you need channels?
		* Channels needed when there is need of communication of goroutines
	How can you send data into a channel?
		* by using `<-` operator i.e. ch <- data
	How can you read data from a channel?
		* by using `<-` operator i.e. data <- ch
*/

/*
	OUTPUT:
	=======
	raja@raja-Latitude-3460:~/Documents/coding/golang/practice-go$ go run goroutines_channels/channels.go
	2021/06/06 10:02:16 [Go-routine] sending reply to main-routine
	2021/06/06 10:02:16 [main-routine] Go-Routine communicated : Hello, Raja
*/
