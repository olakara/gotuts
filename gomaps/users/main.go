package main

import (
	"fmt"
	"github.com/google/uuid"
)

type User struct {
	ID   uuid.UUID
	Name string
	Age  int
}

func main() {
	users := map[uuid.UUID]User{}
	id := uuid.New()

	users[id] = User{
		ID:   id,
		Name: "John Doe",
		Age:  30,
	}

	id = uuid.New()

	users[id] = User{
		ID:   id,
		Name: "Sara Smith",
		Age:  26,
	}

	fmt.Println(users)
}
