package main

import (
	"log"
	"os"
)

func main() {
	fileName := "./files/current_name.txt"
	newName := "./files/new_name.txt"
	removeFile(newName)
	createFile(fileName)
	renameFile(fileName, newName)
}

func createFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Printf("Failed to Create File : %v\n", err)
	}
	defer file.Close()
	log.Printf("Successfully Created file file: %v, %p", fileName, *file)

}

func renameFile(oldName string, newName string) {
	err := os.Rename(oldName, newName)
	if err != nil {
		log.Printf("Failed to Rename : %v\n", err)
	}
	log.Printf("Successfully Renamed file from : %v to : %v", oldName, newName)
}

func removeFile(fileName string) {
	err := os.Remove(fileName)
	if err != nil {
		log.Printf("Failed to Deletefile : %v\n", err)
	}
	log.Printf("Successfully deleted file: %v", fileName)
}

/*
	Rename file:
	============
		Which package has the rename function?

	TO RUN:
	=======
	raja@raja-Latitude-3460:~/Documents/coding/golang/practice-go$ go run files/rename_file.go
	2021/06/03 09:54:34 Failed to Deletefile : remove ./files/new_name.txt: no such file or directory
	2021/06/03 09:54:34 Successfully deleted file: ./files/new_name.txt
	2021/06/03 09:54:34 Successfully Created file file: ./files/current_name.txt, %!p(os.File={0xc00005e180})
	2021/06/03 09:54:34 Successfully Renamed file from : ./files/current_name.txt to : ./files/new_name.txt
	raja@raja-Latitude-3460:~/Documents/coding/golang/practice-go$
*/
