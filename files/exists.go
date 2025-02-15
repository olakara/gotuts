package main

import ("fmt"
"os")

func main() {

	filename:= os.Args[1]
	fmt.Println("Checking if file exists:", filename)
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Println("File does not exist")
		return
	}
}