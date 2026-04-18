package main

import "fmt"

type EntranceGate struct{}

func NewEntranceGate() *EntranceGate {
	return &EntranceGate{}
}

func (eg *EntranceGate) Enter(parkingBiliding *ParkingBuilding, vehicle *Vehcile) *Ticket {
	ticket := parkingBiliding.Allocate(*vehicle)
	return ticket
}

type ExitGate struct {
	CostComputation CostComputation
}

func NewExitGate(costComputation CostComputation) *ExitGate {
	return &ExitGate{
		CostComputation: costComputation,
	}
}

func (eg *ExitGate) Exit(parkingBuilding *ParkingBuilding, ticket *Ticket, payment Payment) {
	parkingBuilding.Release(ticket)

	amount := eg.CostComputation.Compute(ticket)
	fmt.Printf("Total amount is: %v", amount)

	payment.Pay(amount)
}

type ParkignLot struct {
	ParkingBuilding *ParkingBuilding
	EntranceGate    *EntranceGate
	ExitGate        *ExitGate
}

func NewParkingLot(parkingBuilding *ParkingBuilding, entranceGate *EntranceGate, exitGate *ExitGate) *ParkignLot {
	return &ParkignLot{
		ParkingBuilding: parkingBuilding,
		EntranceGate:    entranceGate,
		ExitGate:        exitGate,
	}
}

func (pl *ParkignLot) VehicleArrive(vehicle *Vehcile) *Ticket {
	ticket := pl.EntranceGate.Enter(pl.ParkingBuilding, vehicle)
	return ticket
}

func (pl *ParkignLot) VehicleExit(ticket *Ticket, payment Payment) {
	pl.ExitGate.Exit(pl.ParkingBuilding, ticket, payment)
}
