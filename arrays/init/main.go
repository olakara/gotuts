package main

import "fmt"

func main() {
	// Array of integers
	numbers := [4]int{10, 20, 30, 40}
	n1 := [5]int{}
	n2 := []int{} // slice of integer

	fmt.Printf("%#v\n", numbers)
	fmt.Printf("%#v\n", n1)
	fmt.Printf("%#v\n", n2)
	n3 := append(n2, 10)
	fmt.Printf("%#v\n", n3)
	fmt.Printf("Length of the array: %d\n", len(numbers))
}
