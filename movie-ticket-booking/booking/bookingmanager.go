package booking

import (
	"errors"
	"fmt"
)

type BookingManager struct {
	seatLockManager *SeatLockManager
}

func NewBookingManager(seatLockManager *SeatLockManager) *BookingManager {
	return &BookingManager{
		seatLockManager: seatLockManager,
	}
}

func (bm *BookingManager) CreateBooking(user *User, show *Show, seats []*Seat, paymentStrategy PaymentStrategy) (*Booking, error) {
	if user == nil || show == nil {
		return nil, errors.New("User and Show is required")
	}

	if len(seats) == 0 {
		return nil, errors.New("At least one seat must be selected")
	}

	err := bm.seatLockManager.LockSeats(show, seats, user.Id)
	if err != nil {
		return nil, fmt.Errorf("Couldn't lock the seats: %v", err)
	}

	totalAmount := show.PricingStartegy.CalculatePrice(seats)
	payment := paymentStrategy.Pay(totalAmount)
	if payment.Status != Success {
		bm.seatLockManager.UnlockSeats(show, seats, user.Id)
		return nil, errors.New("Payment was not successful")
	}

	err = bm.seatLockManager.ConfirmSeats(show, seats, user.Id)
	if err != nil {
		bm.seatLockManager.UnlockSeats(show, seats, user.Id)
	}

	return NewBooking(user, show, seats, totalAmount, payment), nil

}
