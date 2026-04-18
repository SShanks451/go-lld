package state

import (
	"vendingMachine/entity"
	"vendingMachine/enums"
)

type Machine interface {
	Inventory() *entity.Inventory
	Balance() int
	AddBalance(value int)
	SetSelectedItemCode(code string)
	SelectedItem() *entity.Item
	DispenseItem()
	RefundBalance()
	Reset()
	SetState(s VendingMachineState)
}

type VendingMachineState interface {
	InsertCoin(coin enums.Coin)
	SelectItem(code string)
	Dispense()
	Refund()
}
