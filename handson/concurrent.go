package main

import (
	"fmt"
	"time"
)

func SayHello(name string) {
	fmt.Println("Hello", name)
}

func main() {

	fmt.Println("In main")
	go SayHello("Abdel")
	go SayHello("Mohamed")
	time.Sleep(3 * time.Second)
	fmt.Println("End of main")
}
