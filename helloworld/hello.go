package main

import "fmt"

const englishHelloPrefix = "Hello, "
const frenchHelloPrefix = "Bonjour, "
const french = "French"
const english = "English"

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch lang {
	case french:
		prefix = frenchHelloPrefix

	case english:

		prefix = englishHelloPrefix
	default:

		prefix = englishHelloPrefix
	}
	return prefix + name + "!"
}

func main() {
	fmt.Println(Hello("", ""))
	fmt.Println(Hello("Abdel", ""))
}
