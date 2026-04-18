package main

type Player struct {
	name  string
	color Color
}

func NewPlayer(name string, color Color) *Player {
	return &Player{
		name:  name,
		color: color,
	}
}
