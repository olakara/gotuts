package main

import "fmt"

func main() {
	// declaring & initialize a map
	users := map[string]string{
		"abdel": "abdel@gmail.com",
		"jane":  "jane@gmail.com",
	}

	fmt.Println("Initial users:", users)

	// adding a new key-value pair
	users["mario"] = "mario@yahoo.com"

	fmt.Println("Updated users:", users)

	// Accessing and displaying a map value
	email := users["abdel"]
	fmt.Println(email)

	// Deleting a key-value pair
	delete(users, "jane")

	fmt.Println("Updated users:", users)

	test := users["unknown"]
	if test != nil {
		fmt.Println(test)
	}

}
