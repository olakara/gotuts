package main

import "fmt"

type Receiver struct{}

func (r *Receiver) Action(actionMessage string) {
	fmt.Println(actionMessage)
}
