package reservation

type Reservation interface {
	GetReservationDate() string
	CalculateCancellationFee() float64
	Cancel()
	GetStatus() string
}

type ReservationBuilder interface {
	SetVertical(vertical string) ReservationBuilder
	SetReservationDate(reservationDate string) ReservationBuilder
	Build() Reservation
}

type reservationBuilderImpl struct {
	vertical        string
	reservationDate string
}

func (r *reservationBuilderImpl) SetVertical(vertical string) ReservationBuilder {
	r.vertical = vertical
	return r
}

func (r *reservationBuilderImpl) SetReservationDate(reservationDate string) ReservationBuilder {
	r.reservationDate = reservationDate
	return r
}

func (r *reservationBuilderImpl) Build() Reservation {
	vertical := r.vertical
	reservationDate := r.reservationDate

	switch vertical {
	case "flight":
		return FlightReservationImpl{reservationDate: reservationDate}
	case "hotel":
		return HotelReservationImpl{reservationDate: reservationDate}
	default:
		return nil
	}
}

func NewReservationBuilder() ReservationBuilder {
	return &reservationBuilderImpl{}
}
