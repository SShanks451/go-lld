package main

func main() {
	// level 1 parking spot managers
	level1Managers := make(map[VehicleType]*ParkingSpotManager)

	level1Managers[FourWheeler] = NewParkingSpotManager([]*ParkingSpot{
		NewParkingSpot("l1-s1", true),
		NewParkingSpot("l1-s2", true),
		NewParkingSpot("l1-s3", true),
	}, &RandomLookupStrategy{})

	level1Managers[TwoWheeler] = NewParkingSpotManager([]*ParkingSpot{
		NewParkingSpot("l1-s4", true),
		NewParkingSpot("l1-s5", true),
		NewParkingSpot("l1-s6", true),
	}, &RandomLookupStrategy{})

	parkingLevel1 := NewParkingLevel(1, level1Managers)

	// level 2 parking spot managers
	level2Managers := make(map[VehicleType]*ParkingSpotManager)

	level2Managers[FourWheeler] = NewParkingSpotManager([]*ParkingSpot{
		NewParkingSpot("l2-s1", true),
		NewParkingSpot("l2-s2", true),
		NewParkingSpot("l2-s3", true),
	}, &RandomLookupStrategy{})

	level2Managers[TwoWheeler] = NewParkingSpotManager([]*ParkingSpot{
		NewParkingSpot("l2-s4", true),
		NewParkingSpot("l2-s5", true),
		NewParkingSpot("l2-s6", true),
	}, &RandomLookupStrategy{})

	parkingLevel2 := NewParkingLevel(2, level2Managers)

	// parking building
	parkingBuilding := NewParkingBuilding([]*ParkingLevel{
		parkingLevel1,
		parkingLevel2,
	})

	// entrance gate
	entranceGate := NewEntranceGate()

	// cost computation
	costComputation := NewCostComputation(&FixedPricingStrategy{})

	// exit gate
	exitGtae := NewExitGate(*costComputation)

	// parking lot
	parkingLot := NewParkingLot(parkingBuilding, entranceGate, exitGtae)

	// vehicles
	vehcile1 := NewVehicle("honda-123", FourWheeler)
	vehicle2 := NewVehicle("bike-123", TwoWheeler)

	vehicle1Ticket := parkingLot.VehicleArrive(vehcile1)
	vehicle2Ticket := parkingLot.VehicleArrive(vehicle2)

	parkingLot.VehicleExit(vehicle2Ticket, &UpiPayment{})
	parkingLot.VehicleExit(vehicle1Ticket, &CashPayment{})
}
