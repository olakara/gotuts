package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <number>")
		return
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Please provide a valid integer.")
		return
	}

	if isPrime(num) {
		fmt.Println("The number is prime!")
	} else {
		fmt.Println("The number is not prime!")
	}
}
