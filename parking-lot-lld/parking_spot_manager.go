package main

import "sync"

type ParkingSpotManager struct {
	Spots          []*ParkingSpot
	LookupStrategy ParkingSpotLookUpStrategy
	Mu             sync.Mutex
}

func NewParkingSpotManager(spots []*ParkingSpot, lookupStrategy ParkingSpotLookUpStrategy) *ParkingSpotManager {
	return &ParkingSpotManager{
		Spots:          spots,
		LookupStrategy: lookupStrategy,
	}
}

func (s *ParkingSpotManager) HasFreeSpot() bool {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	for _, spot := range s.Spots {
		if spot.IsFree {
			return true
		}
	}

	return false
}

func (s *ParkingSpotManager) Park() *ParkingSpot {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	freeSpot := s.LookupStrategy.SelectSpot(s.Spots)
	if freeSpot != nil {
		freeSpot.OccupySpot()
	}

	return freeSpot
}

func (s *ParkingSpotManager) UnPark(spot *ParkingSpot) {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	spot.ReleaseSpot()
}
