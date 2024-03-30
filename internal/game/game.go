package game

import "fmt"

type Game struct {
	Plateau Plateau
}

type scanner struct{}

func (s *scanner) Scanf(format string, a ...interface{}) (n int, err error) {
	return fmt.Scanf(format, a...)
}

type Scanner interface {
	Scanf(format string, a ...interface{}) (n int, err error)
}

func (g *Game) Input(s Scanner) int {
	var input int
	fmt.Println("Entrez quelque chose:")
	s.Scanf("%i", &input)
	fmt.Println("Vous avez entré:", input)
	return input
}

func NewGame(initSeed, initSize int) Game {
	game := Game{
		Plateau: Plateau{},
	}
	game.Plateau.Init(initSeed, initSize)
	return game
}

func (g *Game) Start() {
	g.Plateau.Print()
}
