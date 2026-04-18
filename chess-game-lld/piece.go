package main

type Piece interface {
	canMove(board *Board, from *Cell, to *Cell) bool
	getColor() Color
}

type BasePiece struct {
	color Color
}

func (bp *BasePiece) getColor() Color {
	return bp.color
}

// ----Pawn------
type Pawn struct {
	BasePiece
}

func NewPawn(color Color) *Pawn {
	return &Pawn{
		BasePiece{
			color: color,
		},
	}
}

func (p *Pawn) canMove(board *Board, from *Cell, to *Cell) bool {
	rowDiff := from.row - from.col
	colDiff := to.row - to.col

	if p.getColor() == Black {
		if rowDiff == 1 && colDiff == 0 && to.piece == nil {
			return true
		}

		if from.row == 1 && rowDiff == 2 && colDiff == 0 && to.piece == nil {
			return true
		}

		if rowDiff == 1 && (colDiff == 1 || colDiff == -1) && to.piece != nil {
			return true
		}
	}

	if p.getColor() == White {
		if rowDiff == -1 && colDiff == 0 && to.piece == nil {
			return true
		}

		if from.row == 6 && rowDiff == -2 && colDiff == 0 && to.piece == nil {
			return true
		}

		if rowDiff == -1 && (colDiff == 1 || colDiff == -1) && to.piece != nil {
			return true
		}
	}

	return false
}

// ---------Rook-----
type Rook struct {
	BasePiece
}

func NewRook(color Color) *Rook {
	return &Rook{
		BasePiece{
			color: color,
		},
	}
}

func (r *Rook) canMove(board *Board, from *Cell, to *Cell) bool {
	if from.col == to.col || from.row == to.row {
		return true
	}

	return false
}

// ------Knight-------
type Knight struct {
	BasePiece
}

func NewKnight(color Color) *Knight {
	return &Knight{
		BasePiece{
			color: color,
		},
	}
}

func (k *Knight) canMove(board *Board, from *Cell, to *Cell) bool {
	rowDiff := from.row - to.row
	colDiff := from.col - to.col

	if (rowDiff == 1 || rowDiff == -1) && (colDiff == 2 || colDiff == -2) {
		return true
	}

	if (rowDiff == 2 || rowDiff == -2) && (colDiff == 1 || colDiff == -1) {
		return true
	}

	return false
}

// ----------Bishop-----
type Bishop struct {
	BasePiece
}

func NewBishop(color Color) *Bishop {
	return &Bishop{
		BasePiece{
			color: color,
		},
	}
}

func (k *Bishop) canMove(board *Board, from *Cell, to *Cell) bool {
	return abs(from.row-to.row) == abs(from.col-to.col)
}

// -------Queen--------
type Queen struct {
	BasePiece
}

func NewQueen(color Color) *Queen {
	return &Queen{
		BasePiece{
			color: color,
		},
	}
}

func (q *Queen) canMove(board *Board, from *Cell, to *Cell) bool {
	rowDiff := from.row - to.row
	colDiff := from.col - to.col

	if rowDiff == 0 || colDiff == 0 || abs(rowDiff) == abs(colDiff) {
		return true
	}

	return false
}

// ------King-----
type King struct {
	BasePiece
}

func NewKing(color Color) *King {
	return &King{
		BasePiece{
			color: color,
		},
	}
}

func (k *King) canMove(board *Board, from *Cell, to *Cell) bool {
	rowDiff := from.row - to.row
	colDiff := from.col - to.col

	if abs(rowDiff) <= 1 && abs(colDiff) <= 1 {
		return true
	}

	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
