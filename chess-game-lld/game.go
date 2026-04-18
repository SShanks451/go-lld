package main

import "fmt"

type Game struct {
	board         *Board
	blackPlayer   *Player
	whitePlayer   *Player
	currentPlayer *Player
}

func NewGame() *Game {
	return &Game{
		board: NewBoard(),
	}
}

func (g *Game) SetPlayers(blackName, whiteName string) {
	g.blackPlayer = NewPlayer(blackName, Black)
	g.whitePlayer = NewPlayer(whiteName, White)
	g.currentPlayer = g.whitePlayer
}

func (g *Game) Start() {
	for !g.IsGameOver() {
		fmt.Println("===Current Player Name=== ", g.currentPlayer.name)
		move, err := g.GetPlayerMove()
		if err != nil {
			fmt.Println(err)
			continue
		}

		err = g.board.MovePiece(move)
		if err != nil {
			fmt.Println(err)
			continue
		}

		g.SwitchTurn()
	}

	g.displayResult()
}

func (g *Game) SwitchTurn() {
	if g.currentPlayer == g.blackPlayer {
		g.currentPlayer = g.whitePlayer
	} else {
		g.currentPlayer = g.blackPlayer
	}
}

func (g *Game) IsGameOver() bool {
	return g.board.IsCheckmate(White) || g.board.IsStalemate(White) || g.board.IsCheckmate(Black) || g.board.IsStalemate(Black)
}

func (g *Game) GetPlayerMove() (*Move, error) {
	var frow, fcol, trow, tcol int
	fmt.Println("Enter from row")
	fmt.Scan(&frow)
	fmt.Println("Enter from column")
	fmt.Scan(&fcol)
	fmt.Println("Enter to row")
	fmt.Scan(&trow)
	fmt.Println("Enter to col")
	fmt.Scan(&tcol)

	piece := g.board.GetPiece(frow, fcol)
	if piece == nil || piece.getColor() != g.currentPlayer.color {
		return nil, ErrInvalidSelection
	}

	return &Move{
		from: g.board.GetCell(frow, fcol),
		to:   g.board.GetCell(trow, tcol),
	}, nil

}

func (g *Game) displayResult() {
	switch {
	case g.board.IsCheckmate(White):
		fmt.Println("Black wins by checkmate!")
	case g.board.IsCheckmate(Black):
		fmt.Println("White wins by checkmate!")
	case g.board.IsStalemate(White) || g.board.IsStalemate(Black):
		fmt.Println("The game ends in a stalemate!")
	}
}
