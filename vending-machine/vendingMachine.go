package vendingmachine

import (
	"fmt"
	"sync"
	"vendingMachine/entity"
	"vendingMachine/enums"
	"vendingMachine/state"
)

type VendingMachine struct {
	mu               sync.Mutex
	inventory        *entity.Inventory
	currentState     state.VendingMachineState
	selectedItemCode string
	balance          int
}

var (
	instance *VendingMachine
	once     sync.Once
)

func GetInstance() *VendingMachine {
	once.Do(func() {
		vendingMachine := &VendingMachine{
			inventory: entity.NewInventory(),
		}
		vendingMachine.currentState = state.NewIdleState(vendingMachine)
		instance = vendingMachine
	})

	return instance
}

func (vm *VendingMachine) InsertCoin(coin enums.Coin) {
	vm.mu.Lock()
	defer vm.mu.Unlock()

	vm.currentState.InsertCoin(coin)
}

func (vm *VendingMachine) SelectItem(code string) {
	vm.mu.Lock()
	defer vm.mu.Unlock()

	vm.currentState.SelectItem(code)
}

func (vm *VendingMachine) Dispense() {
	vm.mu.Lock()
	defer vm.mu.Unlock()

	vm.currentState.Dispense()
}

func (vm *VendingMachine) Refund() {
	vm.mu.Lock()
	defer vm.mu.Unlock()

	vm.currentState.Refund()
}

func (vm *VendingMachine) AddItem(code, name string, price, quantity int) *entity.Item {
	item := entity.NewItem(name, code, price)
	vm.inventory.AddItem(code, item, quantity)

	return item
}

// Methods implementing 'machine' interface

func (vm *VendingMachine) Inventory() *entity.Inventory {
	return vm.inventory
}

func (vm *VendingMachine) Balance() int {
	return vm.balance
}

func (vm *VendingMachine) AddBalance(value int) {
	vm.balance += value
}

func (vm *VendingMachine) SetSelectedItemCode(code string) {
	vm.selectedItemCode = code
}

func (vm *VendingMachine) SelectedItem() *entity.Item {
	item := vm.inventory.GetItem(vm.selectedItemCode)
	return item
}

func (vm *VendingMachine) SetState(s state.VendingMachineState) {
	vm.currentState = s
}

func (vm *VendingMachine) Reset() {
	vm.selectedItemCode = ""
	vm.balance = 0
}

func (vm *VendingMachine) RefundBalance() {
	fmt.Println("Refunding amount: ", vm.balance)
	vm.balance = 0
}

func (vm *VendingMachine) DispenseItem() {
	isAvailable := vm.inventory.IsAvailable(vm.selectedItemCode)
	if !isAvailable {
		fmt.Println("Selected item is not avalaible right now...")
		vm.Reset()
		vm.SetState(state.NewIdleState(vm))
		return
	}

	item := vm.inventory.GetItem(vm.selectedItemCode)
	if item == nil {
		fmt.Println("Selected item is invalid...")
		return
	}

	if vm.balance >= item.GetPrice() {
		vm.inventory.ReduceStock(vm.selectedItemCode)
		vm.balance -= item.GetPrice()

		fmt.Println("Item Dispensed...")

		if vm.balance > 0 {
			fmt.Println("Returning change: ", vm.balance)
		}
	}

	vm.Reset()
	vm.SetState(state.NewIdleState(vm))
}
