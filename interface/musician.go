package main

type Musician struct {
	Name string
}

func (m Musician) Perform() {
	println(m.Name, "is performing")
}
