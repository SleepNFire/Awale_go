package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func NewPlateau(seed, size int) Board {
	p := Board{}
	p.Init(seed, size)
	return p
}

type fakeScanner struct {
	inputs []string
	call   int
}

func (f *fakeScanner) Scanf(format string, a ...interface{}) (n int, err error) {
	*(a[0].(*string)) = f.inputs[f.call]
	f.call++
	return 1, nil
}

func NewFakeScanner(inputs []string) fakeScanner {
	return fakeScanner{
		inputs: inputs,
	}
}

func Test_Game(t *testing.T) {

	tests := []struct {
		name  string
		seeds int
		size  int

		fakeS      fakeScanner
		callNumber int
		want       Game
	}{
		{
			name:       "init 1 2",
			seeds:      1,
			size:       2,
			fakeS:      NewFakeScanner([]string{"a", "stop"}),
			callNumber: 2,
			want: Game{
				FinalPlateau: Board{
					north: []int{0, 2},
					south: []int{1, 1},
					info:  InfoBoard{totalSeed: 2 * 2 * 1},
				},
				Move: []int{0},
			},
		},
		{
			name:       "prise",
			seeds:      1,
			size:       2,
			fakeS:      NewFakeScanner([]string{"b", "stop"}),
			callNumber: 2,
			want: Game{
				FinalPlateau: Board{
					north: []int{1, 0},
					south: []int{0, 1},
					info: InfoBoard{
						scoreNorth: 2,
						scoreSouth: 0,
						totalSeed:  2 * 2 * 1,
					},
				},
				Move: []int{1},
			},
		},
		{
			name:       "non prise",
			seeds:      4,
			size:       4,
			fakeS:      NewFakeScanner([]string{"d", "stop"}),
			callNumber: 2,
			want: Game{
				FinalPlateau: Board{
					north: []int{4, 4, 4, 0},
					south: []int{5, 5, 5, 5},
					info: InfoBoard{
						scoreNorth: 0,
						scoreSouth: 0,
						totalSeed:  2 * 4 * 4,
					},
				},
				Move: []int{3},
			},
		},
		{
			name:       "deux tours",
			seeds:      4,
			size:       4,
			fakeS:      NewFakeScanner([]string{"d", "d", "stop"}),
			callNumber: 3,
			want: Game{
				FinalPlateau: Board{
					north: []int{5, 5, 5, 1},
					south: []int{6, 5, 5, 0},
					info: InfoBoard{
						isNorthTurn: true,
						scoreNorth:  0,
						scoreSouth:  0,
						totalSeed:   2 * 4 * 4,
					},
				},
				Move: []int{3, 3},
			},
		},
		{
			name:       "on bad input, second correct",
			seeds:      3,
			size:       3,
			fakeS:      NewFakeScanner([]string{"d", "c", "stop"}),
			callNumber: 3,
			want: Game{
				FinalPlateau: Board{
					north: []int{3, 3, 0},
					south: []int{4, 4, 4},
					info: InfoBoard{
						isNorthTurn: false,
						scoreNorth:  0,
						scoreSouth:  0,
						totalSeed:   2 * 3 * 3,
					},
				},
				Move: []int{2},
			},
		},
		{
			name:       "game over no move",
			seeds:      1,
			size:       1,
			fakeS:      NewFakeScanner([]string{"a"}),
			callNumber: 1,
			want: Game{
				FinalPlateau: Board{
					north: []int{0},
					south: []int{0},
					info: InfoBoard{
						isNorthTurn: false,
						scoreNorth:  2,
						scoreSouth:  0,
						totalSeed:   2 * 1 * 1,
					},
				},
				Move:   []int{0},
				isOver: true,
			},
		},
		{
			name:       "move empty hole",
			seeds:      1,
			size:       3,
			fakeS:      NewFakeScanner([]string{"c", "a", "b", "stop"}),
			callNumber: 4,
			want: Game{
				FinalPlateau: Board{
					north: []int{1, 1, 0},
					south: []int{0, 0, 2},
					info: InfoBoard{
						isNorthTurn: true,
						scoreNorth:  2,
						scoreSouth:  0,
						totalSeed:   2 * 3 * 1,
					},
				},
				Move: []int{2, 1},
			},
		},
		{
			name:       "have most of seed",
			seeds:      2,
			size:       3,
			fakeS:      NewFakeScanner([]string{"b", "c", "c", "c", "a", "a", "b", "b", "c"}),
			callNumber: 9,
			want: Game{
				FinalPlateau: Board{
					north: []int{0, 0, 0},
					south: []int{0, 0, 1},
					info: InfoBoard{
						isNorthTurn: false,
						scoreNorth:  8,
						scoreSouth:  3,
						totalSeed:   2 * 3 * 2,
					},
				},
				Move:   []int{1, 2, 2, 2, 0, 0, 1, 1, 2},
				isOver: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := NewGame(tt.seeds, tt.size)
			assert.Equal(t, NewPlateau(tt.seeds, tt.size), game.FinalPlateau)

			game.s = &tt.fakeS

			game.Start()
			assert.Equal(t, tt.callNumber, tt.fakeS.call)
			assert.Equal(t, tt.want.FinalPlateau, game.FinalPlateau)
			assert.Equal(t, tt.want.Move, game.Move)
		})
	}

}
