package main

import (
	"fmt"
	"time"
)

type person struct {
	Id   int
	GivenName string
	FamilyName string
	StartDate time.Time
	Rating int
}

func main() {

	ids := [3]int{5, 10, 6}
	var name string
	name = "Abdel Raoof"

	println(name)

	fmt.Println(ids)
}
