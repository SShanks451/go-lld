package main

import (
	"fmt"
	"time"
)

type ParkingBuilding struct {
	Levels []*ParkingLevel
}

func NewParkingBuilding(levels []*ParkingLevel) *ParkingBuilding {
	return &ParkingBuilding{
		Levels: levels,
	}
}

func (pb *ParkingBuilding) Allocate(vehicle Vehcile) *Ticket {
	vehicleType := vehicle.VehicleType

	for _, level := range pb.Levels {
		ok := level.HasAvalibility(vehicleType)
		if ok {
			spot := level.Park(vehicleType)
			if spot != nil {
				fmt.Printf("Parking allocated at level: %v and spot: %v", level.LevelNumber, spot.SpotId)
				return &Ticket{
					TicketId:     "ticket-1",
					Vehcile:      &vehicle,
					ParkingLevel: level,
					ParkingSpot:  spot,
					EntryTime:    time.Now(),
				}
			}
		}
	}

	fmt.Println("Parking is full")
	return nil
}

func (pb *ParkingBuilding) Release(ticket *Ticket) {
	ticket.ParkingLevel.UnPark(ticket.Vehcile.VehicleType, ticket.ParkingSpot)
}
