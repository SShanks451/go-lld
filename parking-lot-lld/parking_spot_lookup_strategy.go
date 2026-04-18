package main

type ParkingSpotLookUpStrategy interface {
	SelectSpot(spots []*ParkingSpot) *ParkingSpot
}

type RandomLookupStrategy struct{}
type NearestTopEntryGateLookupStrategy struct{}

func (r *RandomLookupStrategy) SelectSpot(spots []*ParkingSpot) *ParkingSpot {
	for _, spot := range spots {
		if spot.IsFree {
			return spot
		}
	}

	return nil
}
