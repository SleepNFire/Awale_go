package game

import "fmt"

type Board struct {
	north []int
	south []int
	info  InfoBoard
}

type InfoBoard struct {
	totalSeed   int
	isNorthTurn bool
	scoreNorth  int
	scoreSouth  int
}

func (p *Board) Init(seed, size int) {
	p.info.isNorthTurn = true
	for i := 0; i < size; i++ {
		p.north = append(p.north, seed)
		p.south = append(p.south, seed)
		p.info.totalSeed += (seed * 2)
	}
}

func (p *Board) Print() {
	for i := len(p.north) - 1; i >= 0; i-- {
		fmt.Print(" ", p.north[i])
	}
	fmt.Println(" | north score : ", p.info.scoreNorth)
	for i := 0; i < len(p.south); i++ {
		fmt.Print(" ", p.south[i])
	}
	fmt.Println(" | north score : ", p.info.scoreSouth)
}

func playerMove(hole int, ally, opponent []int) int {
	score := 0
	seed := ally[hole]
	ally[hole] = 0
	for i := 1; i < seed+1; i++ {
		position := (hole + i) % len(ally)
		side := ((hole + i) / len(ally)) % 2
		if side == 1 {
			opponent[position]++
			if opponent[position] == 2 || opponent[position] == 3 {
				score += opponent[position]
				opponent[position] = 0
			}
		} else {
			ally[position]++
		}
	}
	return score
}

func (p *Board) Move(hole int) {
	if p.info.isNorthTurn {
		p.info.scoreNorth += playerMove(hole, p.north, p.south)
	} else {
		p.info.scoreSouth += playerMove(hole, p.south, p.north)
	}
	p.info.isNorthTurn = !p.info.isNorthTurn
}

func (p *Board) NotValidMove(hole int, ally []int) bool {
	return ally[hole] == 0
}

func noMove(ally []int) bool {
	for i := 0; i < len(ally); i++ {
		if ally[i] > 0 {
			return false
		}
	}
	return true
}

func (p *Board) WinByScore() bool {
	return (p.info.scoreNorth > (p.info.totalSeed / 2)) || (p.info.scoreSouth > (p.info.totalSeed / 2))
}

func (p *Board) NoMove() bool {
	if p.info.isNorthTurn {
		return noMove(p.north)
	}
	return noMove(p.south)
}
