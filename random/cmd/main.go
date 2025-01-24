package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("A random number: ", rand.Int())
	fmt.Println("Another random number: ", rand.Intn(10))
	fmt.Println("Yet another random number: ", rand.Uint32())
}
