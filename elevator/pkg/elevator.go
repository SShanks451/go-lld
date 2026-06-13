package pkg

import "sync"

type Elevator struct {
	id           int
	state        State
	direction    Direction
	currentFloor int
	requests     []*Request
	observers    []*Observer
	seen         map[Request]bool
	mu           sync.RWMutex
}

func NewElevator(id int) *Elevator {
	return &Elevator{
		id:           id,
		state:        StateIdle,
		direction:    DirectionIdle,
		currentFloor: 1,
	}
}

func (e *Elevator) SetDirection(d Direction) {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.direction = d
}

func (e *Elevator) SetState(s State) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.state == s {
		return
	}
	e.state = s
}

func (e *Elevator) GetRequests() []*Request {
	e.mu.Lock()
	requests := make([]*Request, len(e.requests))
	copy(requests, e.requests)
	e.mu.Unlock()

	return requests
}

func (e *Elevator) GetCurrentFloor() int {
	e.mu.Lock()
	defer e.mu.Unlock()

	return e.currentFloor
}

func (e *Elevator) GetDirection() Direction {
	e.mu.Lock()
	defer e.mu.Unlock()

	return e.direction
}

func (e *Elevator) AddRequest(req *Request) {
	if req == nil {
		return
	}

	e.mu.Lock()
	defer e.mu.Unlock()

	_, ok := e.seen[*req]
	if ok {
		return
	}

	e.requests = append(e.requests, req)
	e.seen[*req] = true

	// if state is idle, start the lift
	if e.state == StateIdle && len(e.requests) > 0 {
		if req.floor > e.currentFloor {
			e.direction = DirectionUp
		} else if req.floor < e.currentFloor {
			e.direction = DirectionDown
		}

		e.state = StateMoving
	}
}

func (e *Elevator) MoveToNextStop(target int) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.state != StateMoving {
		return
	}

	for e.currentFloor != target {
		switch e.direction {
		case DirectionUp:
			e.currentFloor++
		case DirectionDown:
			e.currentFloor--
		default:
			return
		}
	}

	e.completeArrivalLocked()
}

func (e *Elevator) completeArrivalLocked() {
	e.state = StateStopped

	newReq := e.requests[:0]

	for _, r := range e.requests {
		if r.floor == e.currentFloor {
			delete(e.seen, *r)
			continue
		}
		newReq = append(newReq, r)
	}
	e.requests = newReq

	if len(e.requests) == 0 {
		e.state = StateIdle
		e.direction = DirectionIdle
		return
	}

	e.state = StateMoving
}
