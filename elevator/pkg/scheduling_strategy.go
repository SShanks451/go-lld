package pkg

type SchedulingStrategy interface {
	Name() string
	NextStop(e *Elevator) int
}
