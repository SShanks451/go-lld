package main

import (
	"errors"
)

type SnakeGame struct {
	gameBoard         GameBoard
	movementStrategy  MovementStrategy
	snakePositions    []Pair
	snakePositionsMap map[Pair]bool
	foodPositions     []FoodItem
	currentFoodIndex  int
}

func NewSnakeGame(gameBoard GameBoard, foodPositions []FoodItem) *SnakeGame {
	snakePositions := make([]Pair, 0)
	snakePositionsMap := make(map[Pair]bool)

	snakePositions = append(snakePositions, Pair{row: 0, col: 0})
	snakePositionsMap[Pair{row: 0, col: 0}] = true

	return &SnakeGame{
		gameBoard:         gameBoard,
		movementStrategy:  NewHumanMovementStrategy(), // default movement strategy
		snakePositions:    snakePositions,
		snakePositionsMap: snakePositionsMap,
		foodPositions:     foodPositions,
	}
}

func (sg *SnakeGame) SetMovementStrategy(movementStrategy MovementStrategy) {
	sg.movementStrategy = movementStrategy
}

func (sg *SnakeGame) Move(direction string) (int, error) {
	gbWidth := sg.gameBoard.GetWidth()
	gbHeight := sg.gameBoard.GetHeight()

	newPos := sg.movementStrategy.getNextMove(sg.snakePositions[len(sg.snakePositions)-1], direction)
	if newPos.row < 0 || newPos.row >= gbWidth || newPos.col < 0 || newPos.col >= gbHeight {
		return len(sg.snakePositions), errors.New("You hit the wall!! Game Over")
	}

	_, ok := sg.snakePositionsMap[newPos]
	if ok {
		return len(sg.snakePositions), errors.New("You hit yourself!! Game Over")
	}

	sg.snakePositions = append(sg.snakePositions, newPos)
	sg.snakePositionsMap[newPos] = true
	if sg.currentFoodIndex < len(sg.foodPositions) && sg.foodPositions[sg.currentFoodIndex].getRow() == newPos.row && sg.foodPositions[sg.currentFoodIndex].getColumn() == newPos.col {
		sg.currentFoodIndex++
	} else {
		delete(sg.snakePositionsMap, sg.snakePositions[0])
		sg.snakePositions = sg.snakePositions[1:]
	}

	return len(sg.snakePositions), nil

}
