package game

import "fmt"

type Plateau struct {
	north []int
	south []int
}

func (p *Plateau) Init(seed, size int) {
	for i := 0; i < size; i++ {
		p.north = append(p.north, seed)
		p.south = append(p.south, seed)
	}
}

func (p *Plateau) Print() {
	for i := 0; i < len(p.north); i++ {
		fmt.Print(" ", p.north[i])
	}
	fmt.Println()
	for i := len(p.south) - 1; i >= 0; i-- {
		fmt.Print(" ", p.south[i])
	}
	fmt.Println()
}
