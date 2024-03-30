package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func NewPlateau(seed, size int) Plateau {
	p := Plateau{}
	p.Init(seed, size)
	return p
}

type fakeScanner struct {
	input int
}

func (f *fakeScanner) Scanf(format string, a ...interface{}) (n int, err error) {
	*(a[0].(*int)) = f.input
	return 1, nil
}

func Test_Game(t *testing.T) {
	fakeS := fakeScanner{}
	game := NewGame(0, 0)

	assert.Equal(t, NewPlateau(0, 0), game.Plateau)

	game = NewGame(1, 2)
	assert.Equal(t, NewPlateau(1, 2), game.Plateau)

	fakeS.input = 1
	hole := game.Input(&fakeS)
	assert.Equal(t, fakeS.input, hole)

	fakeS.input = 2
	hole = game.Input(&fakeS)
	assert.Equal(t, fakeS.input, hole)
}
