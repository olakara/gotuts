package factory

type Reservation interface {
	GetReservationDate() string
	CalculateCancellationFee() float64
	Cancel()
	GetStatus() string
}

func NewReservation(vertical, reservationDate string) Reservation {
	switch vertical {
	case "flight":
		return FlightReservationImpl{reservationDate: reservationDate}
	case "hotel":
		return HotelReservationImpl{reservationDate: reservationDate}
	default:
		return nil
	}
}
