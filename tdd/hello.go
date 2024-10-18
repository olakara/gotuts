package main

import "fmt"

func Hello() string {
	return "Hello, World!"
}

func HelloName(name string) string {
	return "Hello, " + name + "!"
}

func main() {
	fmt.Println(Hello())
	fmt.Println(HelloName("Abdel"))
}
