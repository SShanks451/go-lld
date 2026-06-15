package booking

type Screen struct {
	Id    string
	Seats []*Seat
}

func NewScreen() *Screen {
	return &Screen{}
}

func (s *Screen) AddSeat(seat *Seat) {
	s.Seats = append(s.Seats, seat)
}
