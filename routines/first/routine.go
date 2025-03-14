package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("main function start")
	go count()
	time.Sleep(20 * time.Millisecond)
	fmt.Println("main function")
	time.Sleep(60 * time.Millisecond)
	fmt.Println("main function end")

}

func count() {
	fmt.Println("count function start")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(2 * time.Millisecond)
	}
	fmt.Println("count function end")
}