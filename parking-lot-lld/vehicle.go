package main

type Vehcile struct {
	VehicleNumber string
	VehicleType   VehicleType
}

func NewVehicle(vehicleNumber string, vehicleType VehicleType) *Vehcile {
	return &Vehcile{
		VehicleNumber: vehicleNumber,
		VehicleType:   vehicleType,
	}
}
