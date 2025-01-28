package main

import (
	"log"
	"os"
)

func main() {

	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	// read the file
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%d bytes: %s\n", n1, string(b1[:n1]))
	// close the file
	defer f.Close()
}
