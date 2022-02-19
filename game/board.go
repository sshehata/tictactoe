package game

import (
  "fmt"
)

// Board a 3 by 3 tictactoe grid 
type board [3][3]tile

// NewBoard create a new tictactoe board
func newBoard() *board {
  return &board{
    {0, 0, 0},
    {0, 0, 0},
    {0, 0, 0},
  }
}

func (b *board) play(t tile, position position) error {
  if b[position.x][position.y] != 0 {
    return fmt.Errorf(fmt.Sprintf("position %v already filled", position))
  }

  b[position.x][position.y] = t
  return nil
}

// Reset reset the board to start a new game
func (b *board) reset() *board {
  return &board{
    {0, 0, 0},
    {0, 0, 0},
    {0, 0, 0},
  }
}

// Moves get all available moves
func (b *board) Moves() []position {
  var moves []position 
  for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
      if b[i][j] == 0 {
        moves = append(moves, position{i, j})
      }
    }
  }

  return moves
}

func (b *board) view() {
  fmt.Println("  _ _ _  ")
  fmt.Println("")
  for i := 0; i < 3; i++ {
    fmt.Printf("| %v %v %v |\n", b[i][0], b[i][1], b[i][2])
  }
  fmt.Println("  _ _ _  ")
}
