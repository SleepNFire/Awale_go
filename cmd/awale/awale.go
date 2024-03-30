package main

import "awale/internal/game"

func main() {
	game := game.NewGame(4, 6)
	game.Start()
}
