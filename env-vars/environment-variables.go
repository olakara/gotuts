package main

import (
	"fmt"
	"os"
)

func main() {

	// Get environment variables
	os.Getenv("KEY")
	fmt.Println("KEY:", os.Getenv("KEY"))
	os.Setenv("KEY", "VALUE")
	fmt.Println("KEY:", os.Getenv("KEY"))
	os.Unsetenv("KEY")
	fmt.Println("KEY:", os.Getenv("KEY"))
}
