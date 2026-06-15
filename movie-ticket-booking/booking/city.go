package booking

import "github.com/google/uuid"

type City struct {
	Id   string
	Name string
}

func NewCity(name string) *City {
	return &City{
		Id:   uuid.NewString(),
		Name: name,
	}
}
