package main

import "fmt"

type Payment interface {
	Pay(amount float64) bool
}

type UpiPayment struct{}

func (up *UpiPayment) Pay(amount float64) bool {
	fmt.Printf("Made UPI Payment of amount: %v", amount)
	return true
}

type CashPayment struct{}

func (cp *CashPayment) Pay(amount float64) bool {
	fmt.Printf("Made Cash Payment of amount: %v", amount)
	return true
}
