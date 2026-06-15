package booking

import (
	"time"

	"github.com/google/uuid"
)

type PaymentStrategy interface {
	Pay(amount float64) *Payment
}

type CreditCardPayment struct {
	CardNumber string
	CVV        int
}

func NewCreditCardPayment(cardNumber string, cvv int) *CreditCardPayment {
	return &CreditCardPayment{
		CardNumber: cardNumber,
		CVV:        cvv,
	}
}

func (cp *CreditCardPayment) Pay(amount float64) *Payment {
	time.Sleep(2 * time.Second)
	return NewPayment(amount, Success, "TXN_"+uuid.NewString())
}
