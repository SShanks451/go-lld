package entities

type Transaction struct {
	from   *User
	to     *User
	amount float64
}

func NewTransaction(from, to *User, amount float64) *Transaction {
	return &Transaction{
		from:   from,
		to:     to,
		amount: amount,
	}
}

func (t *Transaction) GetFrom() *User {
	return t.from
}

func (t *Transaction) GetTo() *User {
	return t.to
}

func (t *Transaction) GetAmount() float64 {
	return t.amount
}
