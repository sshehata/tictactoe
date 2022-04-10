package game

import (
	"fmt"
)

// Position A game board position
type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (p *Position) String() string {
	return fmt.Sprintf("(%v, %v)", p.X, p.Y)
}
