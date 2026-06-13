package entities

import (
	"fmt"
	"maps"
	"math"
	"sync"
)

type BalanceSheet struct {
	owner    *User
	balances map[*User]float64
	mu       sync.Mutex
}

func NewBalanceSheet(owner *User) *BalanceSheet {
	return &BalanceSheet{
		owner:    owner,
		balances: make(map[*User]float64),
	}
}

func (bs *BalanceSheet) GetBalances() map[*User]float64 {
	bs.mu.Lock()
	defer bs.mu.Unlock()

	out := make(map[*User]float64)
	maps.Copy(out, bs.balances)

	return out
}

func (bs *BalanceSheet) AdjustBalance(other *User, amount float64) {
	if other == bs.owner {
		return
	}

	bs.mu.Lock()
	defer bs.mu.Unlock()
	bs.balances[other] += amount
}

func (bs *BalanceSheet) ShowBalances() {
	bs.mu.Lock()
	defer bs.mu.Unlock()

	ownerName := bs.owner.GetName()

	fmt.Printf("--- Balance Sheet for %s ---\n", ownerName)

	var amountOwedToMe, amountIOwe float64

	for k, v := range bs.balances {
		userName := k.GetName()
		if v > 0.01 {
			fmt.Printf("User %v owes amount %v to you.\n", userName, v)
			amountOwedToMe += v
		}
		if v < -0.01 {
			fmt.Printf("You owe amount %v to user %v\n", -v, userName)
			amountIOwe += math.Abs(v)
		}
	}

	fmt.Printf("Total amount owed to you: %v\n", amountOwedToMe)
	fmt.Printf("Total amount you owe: %v\n", amountIOwe)
}
