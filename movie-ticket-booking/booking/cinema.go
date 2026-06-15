package booking

import "github.com/google/uuid"

type Cinema struct {
	Id      string
	Name    string
	City    *City
	Screens []*Screen
}

func NewCinema(name string, city *City, screens []*Screen) *Cinema {
	return &Cinema{
		Id:      uuid.NewString(),
		Name:    name,
		City:    city,
		Screens: screens,
	}
}
