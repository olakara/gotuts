package main

import (
	"fmt"
	"strings"
)

func main() {
	names := []string{"John", "Paul", "George", "Ringo"}
	fmt.Println("Names", names)
	subset := names[1:3]
	fmt.Println("Subset", subset)
	for i, name := range subset {
		subset[i] = strings.ToUpper(name) // -> This will change the original slice
	}
	fmt.Println("Subset ", subset)
	fmt.Println("Names ", names)
}

// convert array to a slice
// slice = array[:]
