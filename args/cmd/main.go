package main

import (
	"fmt"
	"hello"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please provide a name")
		os.Exit(1)
	}
	fmt.Printf(hello.Say(os.Args[1]))

}
