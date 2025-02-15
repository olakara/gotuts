package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Book struct {
	Title         string  `json:"title"`
	Author        string  `json:"author"`
	Price         float32 `json:"price"`
	PublishedDate string  `json:"publishedDate"`
}

func main() {

	// open the data.json file
	fileName := "data.json"
	result, err := os.Stat(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
		panic(err)
	}
	fmt.Println("File Name: ", result.Name())

	// read the data from the file
	data, err := os.ReadFile(fileName)
	fmt.Println("Data: ", string(data))

	var book Book
	err = json.Unmarshal(data, &book)
	if err != nil {
		fmt.Println("Error: ", err)
		panic(err)
	}

	fmt.Printf("Book: %#v\n", book)

}
