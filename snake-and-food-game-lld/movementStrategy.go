package main

type MovementStrategy interface {
	getNextMove(currentHead Pair, direction string) Pair
}

type HumanMovementStrategy struct{}

func NewHumanMovementStrategy() *HumanMovementStrategy {
	return &HumanMovementStrategy{}
}

func (hms *HumanMovementStrategy) getNextMove(currentHead Pair, direction string) Pair {
	switch direction {
	case "U":
		return Pair{currentHead.row - 1, currentHead.col}
	case "D":
		return Pair{currentHead.row + 1, currentHead.col}
	case "L":
		return Pair{currentHead.row, currentHead.col - 1}
	case "R":
		return Pair{currentHead.row, currentHead.col + 1}
	default:
		return Pair{currentHead.row, currentHead.col}
	}
}
