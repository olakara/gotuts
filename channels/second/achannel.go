package main

import (
	"fmt"
	"time"
)

func printCount(c chan int) {
	number := 0
	for number >= 0 {
		number = <-c
		fmt.Print(number, " ")
	}	
}

func main() {
	channel := make(chan int)

	array := []int{8, 10, 5, 6, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

	go printCount(channel)
	for _, number := range array {
		channel <- number
	}
	time.Sleep(time.Second * 2)
	fmt.Println("End of main")

}