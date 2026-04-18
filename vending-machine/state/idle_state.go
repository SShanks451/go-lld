package state

import (
	"fmt"
	"vendingMachine/enums"
)

type IdleState struct {
	machine Machine
}

func NewIdleState(machine Machine) *IdleState {
	return &IdleState{
		machine: machine,
	}
}

func (is *IdleState) InsertCoin(coin enums.Coin) {
	fmt.Println("Select an item first")
}

func (is *IdleState) SelectItem(code string) {
	if !is.machine.Inventory().IsAvailable(code) {
		fmt.Println("Item is not avalaible")
		return
	}

	is.machine.SetSelectedItemCode(code)
	is.machine.SetState(NewItemSelectedState(is.machine))
	fmt.Println("Item Selected: ", code)
}

func (is *IdleState) Dispense() {
	fmt.Println("Select an item first")

}

func (is *IdleState) Refund() {
	fmt.Println("Select an item first")

}
