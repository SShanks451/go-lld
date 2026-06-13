package pkg

type StrategyLook struct{}

func NewStrategyLook() *StrategyLook {
	return &StrategyLook{}
}

func (s *StrategyLook) Name() string {
	return "LOOK"
}

func (s *StrategyLook) NextStop(e *Elevator) int {
	currentFloor := e.GetCurrentFloor()
	currentDirection := e.GetDirection()

	if len(e.requests) == 0 {
		return currentFloor
	}

	nearestFloorAbove, isAbove := 0, false
	nearestFloorBelow, isBelow := 0, false

	for _, req := range e.requests {
		switch {
		case req.GetFloor() > currentFloor:
			if req.GetInternal() || req.GetDirection() == DirectionUp {
				if !isAbove || req.GetFloor() < nearestFloorAbove {
					nearestFloorAbove = req.GetFloor()
					isAbove = true
				}
			}

		case req.floor < currentFloor:
			if req.GetInternal() || req.GetDirection() == DirectionDown {
				if !isBelow || req.GetFloor() > nearestFloorBelow {
					nearestFloorBelow = req.GetFloor()
					isBelow = true
				}
			}
		}
	}

	switch currentDirection {
	case DirectionUp:
		if isAbove {
			return nearestFloorAbove
		} else {
			e.SetDirection(DirectionDown)
			return nearestFloorBelow
		}
	case DirectionDown:
		if isBelow {
			return nearestFloorBelow
		} else {
			e.SetDirection(DirectionUp)
			return nearestFloorAbove
		}
	default:
		switch {
		case isAbove && !isBelow:
			e.SetDirection(DirectionUp)
			return nearestFloorAbove
		case isBelow && !isAbove:
			e.SetDirection(DirectionDown)
			return nearestFloorBelow
		case isAbove && isBelow:
			if abs(currentFloor-nearestFloorAbove) <= abs(currentFloor-nearestFloorBelow) {
				e.SetDirection(DirectionUp)
				return nearestFloorAbove
			} else {
				e.SetDirection(DirectionDown)
				return nearestFloorBelow
			}
		}
	}

	return currentFloor
}

func abs(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}
