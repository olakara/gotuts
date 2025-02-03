package main

import "fmt"

type User struct {
	ID   int
	Name string
	Age  int
}

func main() {

	user := User{
		ID:   1,
		Name: "John",
		Age:  25,
	}

	user2 := User{
		ID:   2,
		Name: "Jane",
		Age:  30,
	}

	users := []User{user, user2}

	fmt.Println(users)
	fmt.Printf("%+v\n", users)
	fmt.Printf("%#v\n", users)
}
