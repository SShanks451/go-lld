package strategy

import (
	"errors"
	"math"
	"splitwise/entities"
)

type ExactSplitStrategy struct{}

func NewExactSplitStrategy() *ExactSplitStrategy {
	return &ExactSplitStrategy{}
}

func (ess *ExactSplitStrategy) CalcualteSplit(total float64, paidBy *entities.User, participants []*entities.User, values []float64) ([]*entities.Split, error) {
	if len(participants) != len(values) {
		return nil, errors.New("Number of participants and split values must be equal")
	}

	if total <= 0.0 {
		return nil, errors.New("Invalid total amount")
	}

	sum := 0.0
	for _, v := range values {
		sum += v
	}

	if math.Abs(total-sum) > 0.01 {
		return nil, errors.New("Sum and final amount is not equal")
	}

	splits := make([]*entities.Split, 0, len(participants))
	for i, p := range participants {
		split := entities.NewSplit(p, values[i])
		splits = append(splits, split)
	}

	return splits, nil
}
