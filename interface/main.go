package main

import "fmt"

func main() {
	musician := Musician{"Mozart"}
	PerformAtVenu(musician)
	fmt.Printf("type %T\n", musician)

	poet := Poet{"William Wordsworth"}
	PerformAtVenu(poet)

}
