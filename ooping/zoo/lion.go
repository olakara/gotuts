package main

type Lion struct {
	Name     string
	location LatLong
}

func (l *Lion) GetLocation() LatLong {
	return l.location
}

func (l *Lion) SetLocation(loc LatLong) {
	l.location = loc
}

func (l *Lion) CanFly() bool {
	return false
}

func (l *Lion) Speak() string {
	return "Roar"
}

func (l *Lion) GetName() string {
	return l.Name
}