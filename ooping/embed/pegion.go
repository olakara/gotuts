package main

import "errors"

type Pegion struct {
	Bird
	Name     string
	location LatLong
	//featherLength float64 // This is overridden from Bird with a different type
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

func (p *Pegion) GetFeatherLength() int {
	return p.featherLength
}

func (p *Pegion) SetFeatherLength(l int) (bool, error) {
	if l < 0 {
		return false, errors.New("feather length cannot be negative")
	}
	p.featherLength = l
	return true, nil
}
