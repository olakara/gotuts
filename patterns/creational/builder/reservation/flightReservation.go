package reservation

type FlightReservation interface {
	Reservation
	AddExtraBaggage()
}

type FlightReservationImpl struct {
	reservationDate string
	status          string
	extraBaggage    bool
}

func (r FlightReservationImpl) GetReservationDate() string {
	return r.reservationDate
}

func (r FlightReservationImpl) CalculateCancellationFee() float64 {
	return 200
}

func (r FlightReservationImpl) Cancel() {
	r.status = "Cancelled"
}

func (r FlightReservationImpl) GetStatus() string {
	return r.status
}
