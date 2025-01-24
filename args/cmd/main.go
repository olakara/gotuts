package main

import (
	"fmt"
	"hello"
	"os"
	"time"
)

func try(str string) {
	fmt.Println(str)
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please provide a name")
		os.Exit(1)
	}

	go try("One")
	go try("Two")
	go try("Three")
	go try("Four")
	go try("Five")
	go try("Six")

	time.Sleep(time.Second * 3)

	fmt.Printf(hello.Say(os.Args[1]))
}
