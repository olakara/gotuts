package main

type Pegion struct {
	Name     string
	location LatLong
}

func (p *Pegion) GetLocation() LatLong {
	return p.location
}

func (p *Pegion) SetLocation(loc LatLong) {
	p.location = loc
}

func (p *Pegion) CanFly() bool {
	return true
}

func (p *Pegion) Speak() string {
	return "Coo"
}

func (p *Pegion) GetName() string {
	return p.Name
}