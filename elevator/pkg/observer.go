package pkg

import "fmt"

type Observer interface {
	OnStateChange(e *Elevator, state State)
	OnFloorChange(e *Elevator, floor int)
}

type ConsoleDisplay struct{}

func NewConsoleDisplay() *ConsoleDisplay {
	return &ConsoleDisplay{}
}

func (c *ConsoleDisplay) OnStateChange(e *Elevator, state State) {
	fmt.Printf("Elevator %v changes state to %v", e.id, state.String())
}

func (c *ConsoleDisplay) OnFloorChange(e *Elevator, floor int) {
	fmt.Printf("Elevator %v changes foor to %v", e.id, floor)
}
