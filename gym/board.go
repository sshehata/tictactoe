package gym

import (
	"fmt"
	"strings"
	"tictactoe/game"
)

// Board A wrapper type around board game type
type Board game.Board

// MakeBoard create a board observation from list
func MakeBoard(list []interface{}) *Board {
	b := new(Board)

	for i, row := range b {
		for j := range row {
			b[i][j] = game.Tile(list[i].([]interface{})[j].(float64))
		}
	}

	return b
}

// Play Modify one position of the board
func (b *Board) Play(t game.Tile, pos game.Position) {
	b[pos.X][pos.Y] = t
}

// Hash Hash current board state into a unique string
func (b *Board) Hash() string {
	var writer strings.Builder
	for _, row := range b {
		for _, col := range row {
			fmt.Fprintf(&writer, "%v", col)
		}
	}

	return writer.String()
}
