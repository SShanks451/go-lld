package booking

import "github.com/google/uuid"

type Seat struct {
	Id         string
	Row        int
	Col        int
	SeatType   SeatType
	SeatStatus SeatStatus
}

func NewSeat(row, col int, seatType SeatType) *Seat {
	return &Seat{
		Id:         uuid.NewString(),
		Row:        row,
		Col:        col,
		SeatType:   seatType,
		SeatStatus: Available,
	}
}
