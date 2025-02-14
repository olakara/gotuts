package main

import (
	"fmt"
	"os"
)

func main() {
	filename := os.Args[1]
	err := os.Remove(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("File deleted successfully")
}
