package main

type Cell struct {
	row   int
	col   int
	piece Piece
}

func NewCell(row, col int) *Cell {
	return &Cell{
		row: row,
		col: col,
	}
}
