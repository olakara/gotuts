package main

import "fmt"

func change(data *int) {
	*data = 10
	fmt.Println("Data:", *data)
}

type Sample struct {
	data int
}

func main() {
	numb := 5
	fmt.Println("Number is:", numb)
	change(&numb)
	fmt.Println("Number is:", numb)
}
