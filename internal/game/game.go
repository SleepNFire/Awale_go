package game

import (
	"awale/internal/tools"
	"fmt"
	"strings"
)

type Game struct {
	s            tools.Scanner
	FinalPlateau Board
	Move         []int
	isOver       bool
}

func NewGame(initSeed, initSize int) Game {
	game := Game{
		FinalPlateau: Board{},
		s:            &tools.Scan{},
		isOver:       false,
	}
	game.FinalPlateau.Init(initSeed, initSize)
	return game
}

func (g *Game) Input() (int, string) {
	var input string

	fmt.Println("Entrez quelque chose:")
	g.s.Scanf("%s", &input)
	fmt.Println("Vous avez entré:", input)

	input = strings.ToLower(input)

	if input == "stop" {
		return -1, "stop"
	}

	if len(input) != 1 || input[0] < 'a' || input[0] > 'z' {
		fmt.Println("Entrée invalide. Veuillez entrer une lettre entre 'a' et 'z'.")
		return g.Input()
	}

	index := int(input[0] - 'a')

	if index < 0 || index > len(g.FinalPlateau.north)-1 {
		return g.Input()
	}

	if g.FinalPlateau.info.isNorthTurn && g.FinalPlateau.NotValidMove(index, g.FinalPlateau.north) {
		return g.Input()
	} else if !g.FinalPlateau.info.isNorthTurn && g.FinalPlateau.NotValidMove(index, g.FinalPlateau.south) {
		return g.Input()
	}

	return index, ""
}

func (g *Game) Finish() {
	g.isOver = g.FinalPlateau.NoMove() || g.FinalPlateau.WinByScore()
}

func (g *Game) Start() {
	turn := 0
	for !g.isOver {
		g.FinalPlateau.Print()
		move, action := g.Input()
		if action == "stop" {
			break
		}
		g.Move = append(g.Move, move)
		g.FinalPlateau.Move(g.Move[turn])
		g.Finish()
		turn++
	}
}
