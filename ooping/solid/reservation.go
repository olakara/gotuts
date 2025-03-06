package main

type Reservation interface {
	GetReservationDate() string
	CalculateCancellationFee() float64
	Cancel()
}

type Customer struct {
	Name string
}

type Seller struct {
	Name string
}
