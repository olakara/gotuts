package main

import (
	"fmt"
	fake "github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"time"
)

type Person struct {
	Id          uuid.UUID
	FirstName   string
	LastName    string
	Gender      string
	DateOfBirth time.Time
}

func NewPerson() (p *Person, err error) {
	return &Person{
		Id:          uuid.New(),
		FirstName:   fake.FirstName(),
		LastName:    fake.LastName(),
		DateOfBirth: fake.Date(),
	}, nil
}

func main() {

	name := fake.FirstName()
	fullName := fake.Person().LastName
	product := fake.Product()

	fmt.Println("Name: ", name)
	fmt.Println("Full Name: ", fullName)
	fmt.Println("Product: ", product)

	person, _ := NewPerson()
	fmt.Println("Person: ", person)
	person.FirstName = "John"
	fmt.Println("Updated Person: ", person)
}
