package pkg

type StrategyFcfs struct{}

func NewStrategyFcfs() *StrategyFcfs {
	return &StrategyFcfs{}
}

func (s *StrategyFcfs) Name() string {
	return "FCFS"
}

func (s *StrategyFcfs) NextStop(e *Elevator) int {
	currentFloor := e.GetCurrentFloor()

	if len(e.requests) == 0 {
		return currentFloor
	}

	requests := e.GetRequests()
	nextFloor := requests[0].GetFloor()

	if nextFloor == currentFloor {
		return currentFloor
	}

	switch e.GetDirection() {
	case DirectionIdle:
		if nextFloor > currentFloor {
			e.SetDirection(DirectionUp)
		}
		if nextFloor < currentFloor {
			e.SetDirection(DirectionDown)
		}
	case DirectionUp:
		if nextFloor < currentFloor {
			e.SetDirection(DirectionDown)
		}
	case DirectionDown:
		if nextFloor > currentFloor {
			e.SetDirection(DirectionUp)
		}
	}

	return nextFloor
}
