package main

import (
	"fmt"
	"os"
)

func main() {

	line := os.Args[1]
	file, err := os.OpenFile("data.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		panic(err)
	}
	file.WriteString(line + "\n")
	defer file.Close()
}
