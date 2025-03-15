package main

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "World", "A name to say hello to.")
var isSpanish bool

func init() {
	flag.BoolVar(&isSpanish, "s", false, "Use Spanish language.")
	flag.BoolVar(&isSpanish, "spanish", false, "Use Spanish language.")
	flag.Parse()
}
func main() {
	if isSpanish {
		fmt.Printf("Hola %s!\n", *name)
	} else {
		fmt.Printf("Hello %s!\n", *name)
	}
}
