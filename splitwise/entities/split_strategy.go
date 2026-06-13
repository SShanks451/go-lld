package entities

type SplitStrategy interface {
	CalcualteSplit(total float64, paidBy *User, participants []*User, values []float64) ([]*Split, error)
}
