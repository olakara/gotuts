package main

type LatLong struct {
	Lat  float64
	Long float64
}

type Animal interface {
	GetLocation() LatLong
	SetLocation(LatLong)
	CanFly() bool
	Speak() string
	GetName() string
}
