package main

type ParkingLevel struct {
	LevelNumber  int
	SpotManagers map[VehicleType]*ParkingSpotManager
}

func NewParkingLevel(levelNumber int, spotManagers map[VehicleType]*ParkingSpotManager) *ParkingLevel {
	return &ParkingLevel{
		LevelNumber:  levelNumber,
		SpotManagers: spotManagers,
	}
}

func (pl *ParkingLevel) HasAvalibility(vehicleType VehicleType) bool {
	manager, ok := pl.SpotManagers[vehicleType]
	if ok {
		return manager.HasFreeSpot()
	}

	return false
}

func (pl *ParkingLevel) Park(vehicleType VehicleType) *ParkingSpot {
	manager, ok := pl.SpotManagers[vehicleType]
	if ok {
		spot := manager.Park()
		return spot
	}

	return nil
}

func (pl *ParkingLevel) UnPark(vehicleType VehicleType, parkingSpot *ParkingSpot) {
	manager, ok := pl.SpotManagers[vehicleType]
	if ok {
		manager.UnPark(parkingSpot)
	}
}
