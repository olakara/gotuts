package main

func say(s string) {
	for i := 0; i < 3; i++ {
		println(s)
	}
}

func main() {
	say("world")
	say("hello")
}
