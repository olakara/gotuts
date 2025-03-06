package main

type Trip struct {
	reservations []Reservation
}

func (t *Trip) CalculateCancellationFee() float64 {
	var total float64
	for _, r := range t.reservations {
		total += r.CalculateCancellationFee()
	}
	return total
}

func (t *Trip) AddReservation(r Reservation) {
	t.reservations = append(t.reservations, r)
}
