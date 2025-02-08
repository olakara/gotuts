package main

type Performer interface {
	Perform()
}

func PerformAtVenu(p Performer) {
	p.Perform()
}
