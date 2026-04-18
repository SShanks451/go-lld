package main

type ParkingSpot struct {
	SpotId string
	IsFree bool
}

func NewParkingSpot(spotId string, isFree bool) *ParkingSpot {
	return &ParkingSpot{
		SpotId: spotId,
		IsFree: isFree,
	}
}

func (s *ParkingSpot) OccupySpot() {
	s.IsFree = false
}

func (s *ParkingSpot) ReleaseSpot() {
	s.IsFree = true
}
