package pkg

import (
	"errors"
	"sync"
)

type Controller struct {
	elevators          []*Elevator
	floors             []*Floor
	schedulingStrategy SchedulingStrategy
	mu                 sync.RWMutex
}

func NewController(numElevator int, numFloor int) *Controller {
	elevators := make([]*Elevator, 0)
	for i := 1; i <= numElevator; i++ {
		elevators = append(elevators, NewElevator(i))
	}

	floors := make([]*Floor, 0)
	for i := 1; i <= numFloor; i++ {
		floors = append(floors, NewFloor(i))
	}

	return &Controller{
		elevators:          elevators,
		floors:             floors,
		schedulingStrategy: NewStrategyFcfs(),
	}
}

func (c *Controller) SetStrategy(s SchedulingStrategy) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.schedulingStrategy = s
}

func (c *Controller) RequestElevator(id, floor int, direction Direction) error {
	elevator := c.getElevatorById(id)
	if elevator == nil {
		return errors.New("No elevator found with this id")
	}

	req := NewRequest(id, floor, direction, false)
	elevator.AddRequest(req)

	return nil
}

func (c *Controller) RequestFloor(id, floor int) error {
	elevator := c.getElevatorById(id)
	if elevator == nil {
		return errors.New("No elevator found with this id")
	}

	direction := DirectionIdle
	if floor > elevator.currentFloor {
		direction = DirectionUp
	} else if floor < elevator.currentFloor {
		direction = DirectionDown
	}

	req := NewRequest(id, floor, direction, true)
	elevator.AddRequest(req)

	return nil
}

func (c *Controller) Step() {
	c.mu.RLock()
	strategy := c.schedulingStrategy
	elevators := make([]*Elevator, len(c.elevators))
	copy(elevators, c.elevators)
	c.mu.RUnlock()

	for _, e := range elevators {
		if len(e.GetRequests()) == 0 {
			continue
		}
		next := strategy.NextStop(e)
		if next != e.GetCurrentFloor() {
			e.MoveToNextStop(next)
		}
	}
}

func (c *Controller) getElevatorById(id int) *Elevator {
	for _, e := range c.elevators {
		if e.id == id {
			return e
		}
	}

	return nil
}
