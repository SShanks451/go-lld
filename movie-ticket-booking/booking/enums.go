package booking

type SeatType uint8

const (
	Regular SeatType = iota
	Premium
	Recliner
)

func (s *SeatType) String() string {
	switch *s {
	case Regular:
		return "Regular"
	case Premium:
		return "Premium"
	case Recliner:
		return "Recliner"
	default:
		return "Regular"
	}
}

func (s *SeatType) BasePrice() float64 {
	switch *s {
	case Regular:
		return 100.0
	case Premium:
		return 200.0
	case Recliner:
		return 300.0
	default:
		return 100.0
	}
}

type SeatStatus uint8

const (
	Available SeatStatus = iota
	Locked
	Booked
)

func (s *SeatStatus) String() string {
	switch *s {
	case Available:
		return "Available"
	case Locked:
		return "Locked"
	case Booked:
		return "Booked"
	default:
		return "Available"
	}
}

type PaymentStatus uint8

const (
	Pending PaymentStatus = iota
	Success
	Failure
)

func (p *PaymentStatus) String() string {
	switch *p {
	case Pending:
		return "Pending"
	case Success:
		return "Success"
	case Failure:
		return "Failure"
	}
	return ""
}
