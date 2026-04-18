package main

import "sync"

type GameBoard struct {
	width  int
	height int
}

var gameBoardInstance *GameBoard
var once sync.Once

func InitGameBoard(width int, height int) {
	once.Do(func() {
		gameBoardInstance = &GameBoard{
			width:  width,
			height: height,
		}
	})
}

func NewGameBoard() *GameBoard {
	if gameBoardInstance == nil {
		panic("GameBoard not initialized")
	}
	return gameBoardInstance
}

func (gb *GameBoard) GetWidth() int {
	return gb.width
}

func (gb *GameBoard) GetHeight() int {
	return gb.height
}
