package main

import "fmt"

func main() {

	var names []string
	names = append(names, "John")
	names = append(names, "Doe", "Smith")
	fmt.Println(names)
	newNames := names
	newNames[0] = "Sara"
	fmt.Println(newNames)

	// merging of two slice to one
	another := append(names, newNames...)
	fmt.Println(another)
	clear(another)
	fmt.Println(another, len(another), cap(another))
	println("Hello World")
}
