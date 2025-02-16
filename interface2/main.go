package main

import "fmt"

func calculateArea(g Geometry) float64 {
	return g.Area()
}

func calculatePerimeter(g Geometry) float64 {
	return g.Perimeter()
}

func describeGeometry(g Geometry) {

	fmt.Println("Area:", calculateArea(g))
	fmt.Println("Perimeter:", calculatePerimeter(g))
}
func main() {

	rect := Rectangle{Width: 10, Height: 5}
	describeGeometry(rect)
	circ := Circle{Radius: 12}
	describeGeometry(circ)
}
