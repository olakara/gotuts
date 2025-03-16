package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Enabled bool
	Path    string
}

func main() {
	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	configuration := Config{}
	err = decoder.Decode(&configuration)
	if err != nil {
		panic(err)
	}
	fmt.Println("Path", configuration.Path)
	fmt.Println("Enabled", configuration.Enabled)
}
