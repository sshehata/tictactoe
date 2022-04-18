package game

import (
	"fmt"
	"strings"
)

// Board a 3 by 3 tictactoe grid
type Board [3][3]Tile

// NewBoard create a new tictactoe board
func newBoard() *Board {
	return &Board{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
}

func (b *Board) play(t Tile, position Position) error {
	if b[position.X][position.Y] != 0 {
		return fmt.Errorf(fmt.Sprintf("position %v already filled", position))
	}

	b[position.X][position.Y] = t
	return nil
}

// Reset reset the board to start a new game
func (b *Board) reset() {
	*b = Board{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
}

// Moves get all available moves
func (b *Board) Moves() []Position {
	var moves []Position
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b[i][j] == 0 {
				moves = append(moves, Position{i, j})
			}
		}
	}

	return moves
}

func (b *Board) String() string {
	var writer strings.Builder

	fmt.Fprintf(&writer, "  _ _ _  \n")
	fmt.Fprintf(&writer, "\n")
	for i := 0; i < 3; i++ {
		fmt.Fprintf(&writer, "| %v %v %v |\n", b[i][0], b[i][1], b[i][2])
	}
	fmt.Fprintf(&writer, "  _ _ _  ")

	return writer.String()
}
