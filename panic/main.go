package main

func init() {
	println("init")
}

func troubledFunction() {

	// recover the panic
	defer func() {
		if r := recover(); r != nil {
			println("recovered from", r)
		}
	}()

	println("troubledFunction")
	panic("another problem")
}

func main() {
	println("main start")
	troubledFunction()
	println("main end")

}
