package main

import "fmt"

type Sample struct {
	data int
}

func change(data Sample) {
	data.data = 10
	fmt.Println("Data:", data)
}

func main() {
	sample := Sample{20}
	fmt.Println("Number is: ", sample)
	change(sample)
	fmt.Println("Number is:", sample)
}
