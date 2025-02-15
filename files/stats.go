package main

import (
	"fmt"
	"os"
)

func main() {

	input := os.Args[1]
	if input == "" {
		fmt.Println("Please provide a file name")
		os.Exit(1)
	}

	file, err := os.Stat(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("File name:", file.Name())
	fmt.Println("Size in bytes:", file.Size())
	fmt.Println("Permissions:", file.Mode())
	fmt.Println("Last modified:", file.ModTime())
	fmt.Println("Is Directory: ", file.IsDir())
	fmt.Println("System interface type: ", file.Sys())
	
}
