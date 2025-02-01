package main

import (
	"io"
	"log"
	"os"
)

func main() {

	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// read the file contents
	data, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%d bytes: %s\n", len(data), data)

}
