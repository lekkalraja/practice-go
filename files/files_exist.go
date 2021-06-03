package main

import (
	"log"
	"os"
	"syscall"
)

func main() {
	checkFileExistsOnLocalDisk("./README.md")
	checkFileExistsOnLocalDisk("./invalid_file.txt")
}

func checkFileExistsOnLocalDisk(fileName string) {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		log.Fatalf("File doesn't exist: %v \n", err)
	}
	log.Printf("File Exist : %v\n", fileInfo)
}

func checkFileOnExternalDisk(fileName string) {
	disk := "/dev/sda"

	fd, err := syscall.Open(disk, syscall.O_RDONLY, 0777)

	if err != nil {
		log.Fatalf("File doesn't exist: %v \n", err)
	}
	log.Printf("File Exist : %v\n", fd)
}

/*
	File exists
	===========
		Check if a file exists on your local disk
		Can you check if a file exists on an external disk?
*/
