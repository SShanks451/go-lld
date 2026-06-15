package booking

import "github.com/google/uuid"

type Payment struct {
	Id            string
	Amount        float64
	Status        PaymentStatus
	TransactionId string
}

func NewPayment(amount float64, status PaymentStatus, transactionId string) *Payment {
	return &Payment{
		Id:            uuid.NewString(),
		Amount:        amount,
		Status:        status,
		TransactionId: transactionId,
	}
}
