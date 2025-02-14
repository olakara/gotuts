package main

import (
	"fmt"
	"os"
)

func main() {

	data := `{ id: 1, name: "John" }`
	file, err := os.OpenFile("test.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		panic(err)
	}
	defer file.Close()

	//file.WriteString(data)
	file.Write([]byte(data))
	fmt.Println("File written successfully")
}
