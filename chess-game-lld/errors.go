package main

import "errors"

var (
	ErrInvalidMove      = errors.New("Invalid Move")
	ErrInvalidSelection = errors.New("Invalid Selection")
	ErrMoveOutOfIndex   = errors.New("Move ot of index")
	ErrorSameColorPiece = errors.New("Same color piece is present on that position")
)
