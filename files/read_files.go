package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	data := readFile("./files/names.txt")
	log.Printf("Got Data :\n%v\n", data)
	names := strings.Split(data, "\n")
	log.Printf("Slice of names : %v\n", names)
}

func readFile(name string) string {
	bs, err := os.ReadFile(name)
	if err != nil {
		log.Fatalf("Something went wrong while reading file : %v\n", err)
	}
	data := string(bs)
	return data
}

/*
	Read file
	=========
		Think of when you’d read a file ‘line by line’ vs ‘at once’?
		 * If you have a very large file, line by line is better because otherwise it won’t fit into memory.
		Create a new file containing names and read it into an array

	For More Ways : https://stackoverflow.com/questions/5884154/read-text-file-into-string-array-and-write
	RUN:
	====
	raja@raja-Latitude-3460:~/Documents/coding/golang/practice-go$ go run files/read_files.go
	2021/06/03 08:55:24 Got Data :
	Janaki
	Ram
	Radha
	Krishna
	2021/06/03 08:55:24 Slice of names : [Janaki Ram Radha Krishna]
*/
