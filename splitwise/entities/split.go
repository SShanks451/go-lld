package entities

type Split struct {
	user   *User
	amount float64
}

func NewSplit(user *User, amount float64) *Split {
	return &Split{
		user:   user,
		amount: amount,
	}
}

func (s *Split) GetUser() *User {
	return s.user
}

func (s *Split) GetAmount() float64 {
	return s.amount
}
