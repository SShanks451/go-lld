package state

import (
	"fmt"
	"vendingMachine/enums"
)

type ItemSelectedState struct {
	machine Machine
}

func NewItemSelectedState(machine Machine) *ItemSelectedState {
	return &ItemSelectedState{
		machine: machine,
	}
}

func (iss *ItemSelectedState) InsertCoin(coin enums.Coin) {
	iss.machine.AddBalance(int(coin))
	fmt.Println("Coin Inserted:", coin)

	if iss.machine.Balance() >= iss.machine.SelectedItem().GetPrice() {
		fmt.Println("Sufficient money received.")
		iss.machine.SetState(NewHasMoneyState(iss.machine))
	}
}

func (iss *ItemSelectedState) SelectItem(code string) {
	fmt.Println("Item already selected")
}

func (iss *ItemSelectedState) Dispense() {
	fmt.Println("Please enter sufficient money.")
}

func (iss *ItemSelectedState) Refund() {
	iss.machine.Reset()
	iss.machine.SetState(NewIdleState(iss.machine))
}
