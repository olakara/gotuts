package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 4, 5
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var round float64 = math.Round(f*100) / 100
	fmt.Println(x, y, f, round)
}
