package main

import "fmt"

func main() {
	listOne := []string{"Apple", "Banana", "Cherry"}
	listTwo := []string{"Water melon", "Kiwi", "Pineapple", "Banana", "Apple", "Pear"}

	fmt.Println("List One: ", listOne, "\nList Two: ", listTwo)
	var combinedList = map[string]int{}
	for _, v := range listOne {
		combinedList[v]++
	}
	for _, v := range listTwo {
		combinedList[v]++
	}

	fmt.Println("Combined List: ", combinedList)

	fmt.Println("Repeated Items: ")
	for v := range combinedList {
		if combinedList[v] > 1 {
			fmt.Println(v)
		}
	}
}
