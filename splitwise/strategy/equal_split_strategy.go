package strategy

import (
	"errors"
	"splitwise/entities"
)

type EqualSplitStrategy struct{}

func NewEqualSplitStrategy() *EqualSplitStrategy {
	return &EqualSplitStrategy{}
}

func (ess *EqualSplitStrategy) CalcualteSplit(total float64, paidBy *entities.User, participants []*entities.User, values []float64) ([]*entities.Split, error) {
	if total <= 0.0 {
		return nil, errors.New("Invalid total amount")
	}

	splits := make([]*entities.Split, 0, len(participants))
	for _, p := range participants {
		split := entities.NewSplit(p, total/float64(len(participants)))
		splits = append(splits, split)
	}

	return splits, nil
}
