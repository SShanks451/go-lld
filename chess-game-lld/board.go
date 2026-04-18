package main

type Board struct {
	cells [8][8]*Cell
}

func NewBoard() *Board {
	b := &Board{}

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			b.cells[i][j] = NewCell(i, j)
		}
	}

	b.InitializeBoard()
	return b
}

func (b *Board) InitializeBoard() {
	for j := 0; j < 8; j++ {
		b.cells[1][j].piece = NewPawn(Black)
		b.cells[6][j].piece = NewPawn(White)
	}

	b.cells[0][0].piece = NewRook(Black)
	b.cells[0][1].piece = NewKnight(Black)
	b.cells[0][2].piece = NewBishop(Black)
	b.cells[0][3].piece = NewQueen(Black)
	b.cells[0][4].piece = NewKing(Black)
	b.cells[0][5].piece = NewBishop(Black)
	b.cells[0][6].piece = NewKnight(Black)
	b.cells[0][7].piece = NewRook(Black)

	b.cells[6][0].piece = NewRook(White)
	b.cells[6][1].piece = NewKnight(White)
	b.cells[6][2].piece = NewBishop(White)
	b.cells[6][3].piece = NewQueen(White)
	b.cells[6][4].piece = NewKing(White)
	b.cells[6][5].piece = NewBishop(White)
	b.cells[6][6].piece = NewKnight(White)
	b.cells[6][7].piece = NewRook(White)
}

func (b *Board) MovePiece(m *Move) error {
	if m.to.row < 0 || m.to.row >= 8 || m.to.col < 0 || m.to.col >= 8 {
		return ErrMoveOutOfIndex
	}

	if m.from.piece == nil {
		return ErrInvalidSelection
	}

	if !m.from.piece.canMove(b, m.from, m.to) {
		return ErrInvalidMove
	}

	if m.to.piece != nil && m.to.piece.getColor() == m.from.piece.getColor() {
		return ErrorSameColorPiece
	}

	piece := m.from.piece
	m.to.piece = piece
	m.from.piece = nil

	return nil
}

func (b *Board) GetPiece(row, col int) Piece {
	if row < 0 || row >= 8 || col < 0 || col >= 8 {
		return nil
	}

	return b.cells[row][col].piece
}

func (b *Board) GetCell(row, col int) *Cell {
	if row < 0 || row >= 8 || col < 0 || col >= 8 {
		return nil
	}

	return b.cells[row][col]
}

func (b *Board) IsCheckmate(color Color) bool {
	// TODO
	return false
}

func (b *Board) IsStalemate(color Color) bool {
	// TODO
	return false
}
