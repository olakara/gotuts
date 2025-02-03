package main

func sayHello(names ...string) {
	for _, name := range names {
		println("Hello", name)
	}
}

func main() {
	sayHello("Alice", "Bob", "Charlie")
}
