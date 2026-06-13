package pkg

type Request struct {
	elevatorId int
	floor      int
	direction  Direction
	internal   bool
}

func NewRequest(elevatorId int, floor int, direction Direction, internal bool) *Request {
	return &Request{
		elevatorId: elevatorId,
		floor:      floor,
		direction:  direction,
		internal:   internal,
	}
}

func (r *Request) GetFloor() int {
	return r.floor
}

func (r *Request) GetInternal() bool {
	return r.internal
}

func (r *Request) GetDirection() Direction {
	return r.direction
}
