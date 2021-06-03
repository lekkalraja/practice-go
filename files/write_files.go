package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	fileName := "./files/my_file.txt"
	removeFile(fileName)
	cities := []string{"Vizag", "Hyderabad", "Bangalore", "Chennai"}
	data := strings.Join(cities, ", ")
	writeToFile(fileName, []byte(data))
}

func writeToFile(fileName string, data []byte) {
	err := os.WriteFile(fileName, data, 0666)
	if err != nil {
		log.Fatalf("Failed to write data to file : %v\n", err)
	}
	log.Printf("Successfully wrote data to : %v", fileName)
}

func removeFile(fileName string) {
	err := os.Remove(fileName)
	if err != nil {
		log.Printf("Failed to Deletefile : %v\n", err)
	}
	log.Printf("Successfully deleted file: %v", fileName)
}

func writeToFileAnotherApproach(fileName string, data []string) {
	file, err := os.Create(fileName)

	if err != nil {
		return
	}
	defer file.Close()

	for i := 0; i < len(data); i++ {
		file.WriteString(data[i])
		file.WriteString("\n")
	}
}

/*
	Write file
	==========
		* Write a list of cities to a new file.

	RUN:
	====
	raja@raja-Latitude-3460:~/Documents/coding/golang/practice-go$ go run files/write_files.go
	2021/06/03 09:53:52 Failed to Deletefile : remove ./files/my_file.txt: no such file or directory
	2021/06/03 09:53:52 Successfully deleted file: ./files/my_file.txt
	2021/06/03 09:53:52 Successfully wrote data to : ./files/my_file.txt
	raja@raja-Latitude-3460:~/Documents/coding/golang/practice-go$
*/
