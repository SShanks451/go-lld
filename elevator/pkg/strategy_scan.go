package pkg

type StrategyScan struct{}

func NewStrategyScan() *StrategyScan {
	return &StrategyScan{}
}

func (s *StrategyScan) Name() string {
	return "SCAN"
}

// func (s *StrategyScan) NextStop(e *Elevator) int {

// }
