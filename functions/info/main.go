package main

import "fmt"

func getInfo(names []string) (string, int, int) {
	formattedString := fmt.Sprintf("The names are: %v", names)
	length := len(names)
	capacity := cap(names)
	return formattedString, length, capacity
}

func main() {
	names := []string{"John", "Paul", "George", "Ringo"}

	formattedString, length, capacity := getInfo(names)
	fmt.Println(formattedString)
	fmt.Println("Length:", length)
	fmt.Println("Capacity:", capacity)
}
