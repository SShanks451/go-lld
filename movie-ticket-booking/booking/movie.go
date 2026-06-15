package booking

import (
	"time"

	"github.com/google/uuid"
)

type Movie struct {
	Id       string
	Title    string
	Duration time.Duration
}

func NewMovie(title string, duration time.Duration) *Movie {
	return &Movie{
		Id:       uuid.NewString(),
		Title:    title,
		Duration: duration,
	}
}
