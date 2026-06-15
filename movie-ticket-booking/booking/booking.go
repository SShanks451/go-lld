package booking

import "github.com/google/uuid"

type Booking struct {
	Id          string
	User        *User
	Show        *Show
	Seats       []*Seat
	TotalAmount float64
	Payment     *Payment
}

func NewBooking(user *User, show *Show, seats []*Seat, totalAmount float64, payment *Payment) *Booking {
	return &Booking{
		Id:          uuid.NewString(),
		User:        user,
		Show:        show,
		Seats:       seats,
		TotalAmount: totalAmount,
		Payment:     payment,
	}
}
