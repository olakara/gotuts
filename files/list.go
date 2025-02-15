package main

import (
	"os"
)

func main() {

	files, err := os.ReadDir(".")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		println(file.Name())
	}
}
