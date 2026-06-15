package booking

type PricingStartegy interface {
	CalculatePrice(seats []*Seat) float64
}

type WeekdayPricingStrategy struct{}

func NewWeekdayPricingStrategy() *WeekdayPricingStrategy {
	return &WeekdayPricingStrategy{}
}

func (w *WeekdayPricingStrategy) CalculatePrice(seats []*Seat) float64 {
	total := 0.0
	for _, seat := range seats {
		total += seat.SeatType.BasePrice()
	}

	return total
}

var weekendSurge float64 = 2.0

type WeekendPricingStrategy struct{}

func NewWeekendPricingStrategy() *WeekendPricingStrategy {
	return &WeekendPricingStrategy{}
}

func (w *WeekendPricingStrategy) CalculatePrice(seats []*Seat) float64 {
	total := 0.0
	for _, seat := range seats {
		total += seat.SeatType.BasePrice()
	}

	return total * weekendSurge
}
