package main

import (
	"fmt"
	"os"
)

func main() {

	var data, err = os.ReadFile("data.json")
	if err != nil {
		fmt.Println("Error reading file: ", err)
		panic(err)
	}
	fmt.Println(string(data))
}
