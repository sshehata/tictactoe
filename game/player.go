package game

import (
	"fmt"
)

type player struct {
	rowContainer             [3]int
	columnContainer          [3]int
	diagonalContainer        int
	reverseDiagonalContainer int
	tile                     Tile
}

func newPlayer(t Tile) *player {
	return &player{
		tile: t,
	}
}

func (p *player) checkWinner(x, y int) bool {
	p.rowContainer[x]++
	if p.rowContainer[x] == 3 {
		return true
	}

	p.columnContainer[y]++
	if p.columnContainer[y] == 3 {
		return true
	}

	if x == y {
		p.diagonalContainer++
		if p.diagonalContainer == 3 {
			return true
		}
	}

	if x+y == 2 {
		p.reverseDiagonalContainer++
		if p.reverseDiagonalContainer == 3 {
			return true
		}
	}

	return false
}

func (p *player) reset() {
	p.diagonalContainer = 0
	p.reverseDiagonalContainer = 0
	p.rowContainer = [3]int{0, 0, 0}
	p.columnContainer = [3]int{0, 0, 0}
}

func (p *player) String() string {
	return fmt.Sprintf("%v", p.tile)
}

func (p *player) Tile() Tile {
	return p.tile
}
