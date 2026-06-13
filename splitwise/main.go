package main

import (
	"fmt"
	"splitwise/entities"
	"splitwise/service"
	"splitwise/strategy"
)

func main() {
	// 1. Setup the service ----------------------------------------------------
	svc := service.GetInstance()

	// 2. Create users and a group --------------------------------------------
	alice := svc.AddUser("1", "Alice", "alice@a.com")
	bob := svc.AddUser("2", "Bob", "bob@b.com")
	charlie := svc.AddUser("3", "Charlie", "charlie@c.com")
	// david := svc.AddUser("4", "David", "david@d.com")

	// 3. Use Case 1: Equal Split ---------------------------------------------
	fmt.Println("--- Use Case 1: Equal Split ---")
	err := svc.CreateExpense(entities.NewExpenseBuilder().
		SetId("expense-1").
		SetDescription("Jaipur trip").
		SetAmount(1000.0).
		SetPaidBy(alice).
		SetParticipants([]*entities.User{bob, charlie}).
		SetSplitStrategy(strategy.NewEqualSplitStrategy()),
	)

	if err != nil {
		fmt.Println("error:", err)

	}

	svc.ShowBalanceSheet(alice.GetId())
	svc.ShowBalanceSheet(bob.GetId())
	svc.ShowBalanceSheet(charlie.GetId())
	fmt.Println()

	// 4. Use Case 2: Exact Split ---------------------------------------------
	fmt.Println("--- Use Case 2: Exact Split ---")
	if err := svc.CreateExpense(entities.NewExpenseBuilder().
		SetId("expense-2").
		SetDescription("Movie Tickets").
		SetAmount(500.0).
		SetPaidBy(alice).
		SetParticipants([]*entities.User{bob, charlie}).
		SetSplitStrategy(strategy.NewExactSplitStrategy()).
		SetSplitValues([]float64{200.0, 300.0}),
	); err != nil {
		fmt.Println("error:", err)
	}

	svc.ShowBalanceSheet(alice.GetId())
	svc.ShowBalanceSheet(bob.GetId())
	svc.ShowBalanceSheet(charlie.GetId())
	fmt.Println()

	// 5. Use Case 3: Percentage Split ----------------------------------------
	fmt.Println("--- Use Case 3: Percentage Split ---")
	if err := svc.CreateExpense(entities.NewExpenseBuilder().
		SetId("expense-3").
		SetDescription("Groceries").
		SetAmount(500.0).
		SetPaidBy(bob).
		SetParticipants([]*entities.User{alice, charlie}).
		SetSplitStrategy(strategy.NewPercentageSplitStrategy()).
		SetSplitValues([]float64{40.0, 60.0}), // 40%, 70%
	); err != nil {
		fmt.Println("error:", err)
	}

	svc.ShowBalanceSheet(alice.GetId())
	svc.ShowBalanceSheet(bob.GetId())
	svc.ShowBalanceSheet(charlie.GetId())
	fmt.Println()

	// 7. Use Case 5: Partial Settlement --------------------------------------
	fmt.Println("--- Use Case 5: Partial Settlement ---")
	// From the simplified debts we can see Bob should pay Alice. Bob pays 100.
	_ = svc.SettleUp(bob.GetId(), alice.GetId(), 100)

	fmt.Println("--- Balances After Partial Settlement ---")
	svc.ShowBalanceSheet(alice.GetId())
	svc.ShowBalanceSheet(bob.GetId())
}
