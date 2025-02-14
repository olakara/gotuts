package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Book struct {
	Title         string
	Author        string
	Price         float32
	PublishedDate string
}

func main() {

	now := time.Now()
	formattedTime := now.Format(time.RFC3339)

	book := Book{
		Title:         "Another book",
		Author:        "John Doe",
		Price:         11.95,
		PublishedDate: formattedTime,
	}
	bookJson, err := json.Marshal(book)
	if err != nil {
		fmt.Println("Error: ", err)
		panic(err)
	}
	fmt.Println(string(bookJson))

	file, err := os.OpenFile("data.json", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		panic(err)
	}
	defer file.Close()

	file.WriteString(string(bookJson))

}
