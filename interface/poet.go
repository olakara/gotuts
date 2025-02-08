package main

type Poet struct {
	Name string
}

func (p Poet) Perform() {
	println(p.Name, "is performing")
}
