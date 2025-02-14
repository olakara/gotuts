package main

import ("fmt"
	"os")
	

func main() {

	file, err := os.Create("file.txt")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		return
	}

	defer file.Close()

	fmt.Println("File created successfully")
}