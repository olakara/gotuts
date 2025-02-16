package main

import "fmt"

func main() {

	rect := Rectangle{Width: 10, Height: 5}
	fmt.Println("Rectangle Area: ", rect.Area())
	circ := Circle{Radius: 12}
	fmt.Println("Circle Area: ", circ.Area())
}
