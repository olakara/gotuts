package main

import "fmt"

func main() {

	// Array of four strings
	names := [4]string{"John", "Paul", "George", "Ringo"}
	fmt.Printf("%#v\n", names)
	fmt.Println(names[2])

	// Slice of strings
	sliceNames := []string{"John", "Paul", "George", "Ringo"}
	fmt.Printf("%#v\n", sliceNames)
	fmt.Println(sliceNames[2])
}
