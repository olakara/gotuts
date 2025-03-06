package main

type HotelReservation interface {
	Reservation
	ChangeType()
}
type HotelReservationImpl struct {
	reservationDate string
	status          string
}

func (r HotelReservationImpl) GetReservationDate() string {
	return r.reservationDate
}

func (r HotelReservationImpl) CalculateCancellationFee() float64 {
	return 50
}

func (r HotelReservationImpl) Cancel() {
	r.status = "Cancelled"
}
