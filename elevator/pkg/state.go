package pkg

type State int

const (
	StateIdle State = iota
	StateMoving
	StateStopped
)

func (s *State) String() string {
	switch *s {
	case StateIdle:
		return "IDLE"
	case StateMoving:
		return "MOVING"
	case StateStopped:
		return "STOPPED"
	}

	return ""
}
