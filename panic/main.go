package main

func init() {
	println("init")
}

func troubledFunction() {

	println("troubledFunction")
	panic("another problem")
}

func main() {
	println("main start")
	troubledFunction()
	println("main end")

}
