package main

import (
	"crypto/rand"
	"log"
	"fmt"
	"os"
)

func generateRandomData(size int64) []byte {
	data := make([]byte, size)
	_, err := rand.Read(data)
	if err != nil {
		panic(err)
	}
	return data
}

func Shred(path string) {
	// Get the size of the file
	fileInfo, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	fileSize := fileInfo.Size()

	fmt.Println("\n Shredding begins...")
	// Overwrite the file 3 times with random data
	for i := 0; i < 3; i++ {
		randomData := generateRandomData(fileSize)

		file, err := os.OpenFile(path, os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		_, err = file.WriteAt(randomData, 0)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("\n Overwrite iteration ", i+1)
		}
	}

	// Delete the file
	err = os.Remove(path)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("\n Shredding is complete !!")
	}
}

func main() {

    file, errs := os.Create("randomfile.txt")
    if errs != nil {
        log.Fatalf("failed creating file: %s", errs)
    }
    defer file.Close()

    //write to randomfile
    _, errs = file.WriteString("This is a random file."+
            " This text will be overwritten 3 times"+
                " and then text file shall be deleted.")

    if errs != nil {
        log.Fatalf("failed writing to file: %s", errs)
    } else {
	fmt.Println("\n Created and Wrote to randomfile.txt succesfully !!")
    }
//	now shred the randomfile
    Shred("randomfile.txt")
}
