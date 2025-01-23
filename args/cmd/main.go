package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please provide a name")
		os.Exit(1)
	}
	fmt.Printf("Hello %s\n", os.Args[1])

}
