package entities

import (
	"errors"
	"time"
)

type Expense struct {
	id          string
	description string
	amount      float64
	paidBy      *User
	splits      []*Split
	timestamp   time.Time
}

func (e *Expense) GetId() string {
	return e.id
}

func (e *Expense) GetDescription() string {
	return e.description
}

func (e *Expense) GetAmount() float64 {
	return e.amount
}

func (e *Expense) GetPaidBy() *User {
	return e.paidBy
}

func (e *Expense) GetSplits() []*Split {
	return e.splits
}

func (e *Expense) GetTimestamp() time.Time {
	return e.timestamp
}

type ExpenseBuilder struct {
	id            string
	description   string
	amount        float64
	paidBy        *User
	participants  []*User
	splitStrategy SplitStrategy
	splitValues   []float64
}

func NewExpenseBuilder() *ExpenseBuilder {
	return &ExpenseBuilder{}
}

func (eb *ExpenseBuilder) SetId(id string) *ExpenseBuilder {
	eb.id = id
	return eb
}

func (eb *ExpenseBuilder) SetDescription(d string) *ExpenseBuilder {
	eb.description = d
	return eb
}

func (eb *ExpenseBuilder) SetAmount(a float64) *ExpenseBuilder {
	eb.amount = a
	return eb
}

func (eb *ExpenseBuilder) SetPaidBy(u *User) *ExpenseBuilder {
	eb.paidBy = u
	return eb
}

func (eb *ExpenseBuilder) SetParticipants(participants []*User) *ExpenseBuilder {
	eb.participants = participants
	return eb
}

func (eb *ExpenseBuilder) SetSplitStrategy(s SplitStrategy) *ExpenseBuilder {
	eb.splitStrategy = s
	return eb
}

func (eb *ExpenseBuilder) SetSplitValues(v []float64) *ExpenseBuilder {
	eb.splitValues = v
	return eb
}

func (eb *ExpenseBuilder) Build() (*Expense, error) {
	if eb.paidBy == nil {
		return nil, errors.New("Paid by is required")
	}

	if eb.splitStrategy == nil {
		return nil, errors.New("Split strategy is required")
	}

	if len(eb.participants) == 0 {
		return nil, errors.New("At least 1 participant must be there")
	}

	splits, err := eb.splitStrategy.CalcualteSplit(eb.amount, eb.paidBy, eb.participants, eb.splitValues)
	if err != nil {
		return nil, err
	}

	return &Expense{
		id:          eb.id,
		description: eb.description,
		amount:      eb.amount,
		paidBy:      eb.paidBy,
		splits:      splits,
		timestamp:   time.Now(),
	}, nil
}
