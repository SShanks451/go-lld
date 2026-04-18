package main

import "time"

type Ticket struct {
	TicketId     string
	Vehcile      *Vehcile
	ParkingLevel *ParkingLevel
	ParkingSpot  *ParkingSpot
	EntryTime    time.Time
}
