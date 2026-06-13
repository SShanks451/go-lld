package strategy

import (
	"errors"
	"math"
	"splitwise/entities"
)

type PercentageSplitStrategy struct{}

func NewPercentageSplitStrategy() *PercentageSplitStrategy {
	return &PercentageSplitStrategy{}
}

func (pss *PercentageSplitStrategy) CalcualteSplit(total float64, paidBy *entities.User, participants []*entities.User, values []float64) ([]*entities.Split, error) {
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

	if math.Abs(100-sum) > 0.01 {
		return nil, errors.New("Invalid split percentage")
	}

	splits := make([]*entities.Split, 0, len(participants))
	for i, p := range participants {
		split := entities.NewSplit(p, (values[i]*total)/100.0)
		splits = append(splits, split)
	}

	return splits, nil
}
