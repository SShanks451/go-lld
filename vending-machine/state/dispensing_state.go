package state

import (
	"fmt"
	"vendingMachine/enums"
)

type DispensingState struct {
	machine Machine
}

func NewDispensingState(machine Machine) *DispensingState {
	return &DispensingState{
		machine: machine,
	}
}

func (ds *DispensingState) InsertCoin(coin enums.Coin) {
	fmt.Println("Currently dispensing. Please wait.")
}

func (ds *DispensingState) SelectItem(code string) {
	fmt.Println("Currently dispensing. Please wait.")
}

func (ds *DispensingState) Dispense() {
	// no op
}

func (ds *DispensingState) Refund() {
	fmt.Println("Dispensing in progress. Refund not allowed.")
}
