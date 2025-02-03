package main

import "fmt"

type User struct {
	ID   int
	Name string
	Age  int
}

func (u User) GetID() int {
	return u.ID
}

func (u User) String() string {
	return fmt.Sprintf("ID: %d, Name: %s, Age: %d", u.ID, u.Name, u.Age)
}

func main() {
	u := User{ID: 1, Name: "John Doe", Age: 25}

	fmt.Println(u.GetID())
	fmt.Println(u.String())
}
