package main

import "fmt"

type Pair struct {
	row int
	col int
}

func main() {
	width := 5
	height := 5

	InitGameBoard(width, height)
	gameBoard := NewGameBoard()

	foodPositions := []FoodItem{
		FoodFactory("normal", Pair{2, 0}),
		FoodFactory("normal", Pair{3, 2}),
		FoodFactory("normal", Pair{4, 4}),
	}

	var movementStrategy MovementStrategy = NewHumanMovementStrategy()

	snakeGame := NewSnakeGame(*gameBoard, foodPositions)
	snakeGame.SetMovementStrategy(movementStrategy)

	for {
		for i := range width {
			for j := range height {
				if _, ok := snakeGame.snakePositionsMap[Pair{row: i, col: j}]; ok {
					fmt.Print("*  ")
				} else {
					fmt.Print("_  ")
				}
			}
			fmt.Println()
		}
		fmt.Println()

		fmt.Println("Enter the direction: ")
		var direction string
		fmt.Scanln(&direction)

		if direction != "U" && direction != "D" && direction != "L" && direction != "R" {
			fmt.Println("Invalid Move")
			continue
		}
		score, err := snakeGame.Move(direction)
		fmt.Println("Current Score: ", score)

		if err != nil {
			fmt.Println(err.Error())
			break
		}
	}
}
