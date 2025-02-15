package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Reading file line by line")
	file, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
