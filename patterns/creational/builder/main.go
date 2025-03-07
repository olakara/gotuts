package main

import (
	reservation "builder/reservation"
	"fmt"
)

func main() {
	builder := reservation.NewReservationBuilder()
	flightReservation := builder.SetVertical("flight").SetReservationDate("2021-01-01").Build()
	fmt.Println("flightReservation:", flightReservation)
}
