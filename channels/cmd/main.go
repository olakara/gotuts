package main

func GetMessage(myChannel chan string) {
	myChannel <- "Hello from function"
}

func main() {
	myChannel := make(chan string)
	anotherChannel := make(chan string)

	go GetMessage(myChannel)

	go func() {
		myChannel <- "Hello from channel"
	}()
	go func() {
		anotherChannel <- "Hello from another channel"
	}()

	select {
	case msg := <-myChannel:
		println(msg)
	case msg := <-anotherChannel:
		println(msg)
	}
}
