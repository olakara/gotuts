package main

import "fmt"

type Animal struct {
	Name   string
	canFly bool
}


func main() {

	anAnimal := Animal{Name: "Bird", canFly: true}	
	fmt.Println(anAnimal.Name)
	fmt.Println(anAnimal.canFly)

	fmt.Println("Pointer to the structure")
	aPtr := &anAnimal
	fmt.Println(aPtr.Name)
	fmt.Println(aPtr.canFly)

	fmt.Println("Change the value of the structure")
	anAnimal.Name = "Dog"
	anAnimal.canFly = false
	fmt.Println(anAnimal.Name)
	fmt.Println(anAnimal.canFly)

	fmt.Println("Displaying with Pointer")
	fmt.Println(aPtr.Name)
	fmt.Println(aPtr.canFly)

	fmt.Println("Change the value of the structure using Pointer")
	aPtr.canFly = true

	fmt.Println("Displaying with Pointer")
	fmt.Println(aPtr.Name)
	fmt.Println(aPtr.canFly)

}
