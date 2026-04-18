package main

import (
	"fmt"
	vendingmachine "vendingMachine"
	"vendingMachine/enums"
)

// Demo — mirrors VendingMachineDemo.java step-by-step.
func main() {
	vm := vendingmachine.GetInstance()

	// Stock the machine
	vm.AddItem("A1", "Coke", 25, 10)
	vm.AddItem("A2", "Pepsi", 25, 2)
	vm.AddItem("B1", "Water", 10, 5)

	fmt.Println("\n--- Step 1: Select an item ---")
	vm.SelectItem("A1")

	fmt.Println("\n--- Step 2: Insert coins ---")
	vm.InsertCoin(enums.DIME)   // 10
	vm.InsertCoin(enums.DIME)   // 10
	vm.InsertCoin(enums.NICKEL) // 5

	fmt.Println("\n--- Step 3: Dispense item ---")
	vm.Dispense() // Should dispense Coke

	fmt.Println("\n--- Step 4: Select another item ---")
	vm.SelectItem("B1")

	fmt.Println("\n--- Step 5: Insert more than needed ---")
	vm.InsertCoin(enums.QUARTER) // 25, price is 10 -> change 15

	fmt.Println("\n--- Step 6: Dispense and return change ---")
	vm.Dispense()
}
