package state

import (
	"fmt"
	"vendingMachine/enums"
)

type HasMoneyState struct {
	machine Machine
}

func NewHasMoneyState(machine Machine) *HasMoneyState {
	return &HasMoneyState{
		machine: machine,
	}
}

func (hms *HasMoneyState) InsertCoin(coin enums.Coin) {
	fmt.Println("Sufficient Coin already inserted")
}

func (hms *HasMoneyState) SelectItem(code string) {
	fmt.Println("Item already selected")
}

func (hms *HasMoneyState) Dispense() {
	fmt.Println("Dispensing item: ", hms.machine.SelectedItem().GetName())
	hms.machine.SetState(NewDispensingState(hms.machine))
	hms.machine.DispenseItem()
}

func (hms *HasMoneyState) Refund() {
	hms.machine.RefundBalance()
	hms.machine.Reset()
	hms.machine.SetState(NewIdleState(hms.machine))
}
