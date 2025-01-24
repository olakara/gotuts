package main

import "fmt"

func GetMessage(myChannel chan string) {
	myChannel <- "Hello"
}

func main() {
	myChannel := make(chan string)
	go GetMessage(myChannel)
	message := <-myChannel
	fmt.Println("Message from channel: ", message)
}
