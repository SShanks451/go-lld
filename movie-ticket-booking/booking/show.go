package booking

import (
	"time"

	"github.com/google/uuid"
)

type Show struct {
	Id              string
	Movie           *Movie
	Screen          *Screen
	StartTime       time.Time
	PricingStartegy PricingStartegy
}

func NewShow(movie *Movie, screen *Screen, startTime time.Time, pricingStrategy PricingStartegy) *Show {
	return &Show{
		Id:              uuid.NewString(),
		Movie:           movie,
		Screen:          screen,
		StartTime:       startTime,
		PricingStartegy: pricingStrategy,
	}
}
