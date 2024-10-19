package main

import (
	"fmt"
)

func main() {

	var actualStringVariable string
	var addressOfStringVariable *string
	fmt.Println("actualStringVariable: ", actualStringVariable)
	fmt.Println("addressOfStringVariable: ", addressOfStringVariable)

	var p *string
	p = new(string)
	*p = "Hello"
	fmt.Println("value of p	: ", p)   // will display the address of the string
	fmt.Println("value of *p	: ", *p) // will display the value of the string

}
