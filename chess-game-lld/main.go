package main

func main() {
	game := NewGame()
	game.SetPlayers("black", "white")
	game.Start()
}
