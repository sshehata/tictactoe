package game

import (
  "fmt"
  "errors"

  "tictactoe/render"
)

// Game a game of tictactoe
type Game struct{
  Board *board
  currentPlayer tile
  moveCount int
  gameOver bool
  winner *player
  players [3]*player
}

// NewGame Start a new tictactoe game
func NewGame() *Game {
  return &Game{
    Board: newBoard(),
    currentPlayer: 1,
    players: [3]*player{
      newPlayer(undefined),
      newPlayer(otile),
      newPlayer(xtile),
    },
  }
}

// Play current player makes a move
func (g *Game) Play(x, y int) error {
  if g.gameOver {
    return errors.New("the game is over") 
  }

  p := g.players[g.currentPlayer]
  err := g.Board.play(g.players[g.currentPlayer].tile, position{x, y})
  if err != nil {
    return err
  }

  if p.checkWinner(x, y) {
    g.gameOver = true
    g.winner = p
  }

  g.moveCount++
  if g.moveCount >= 9 {
    g.gameOver = true
  }

  switch g.currentPlayer {
  case otile:
    g.currentPlayer = xtile
  case xtile:
    g.currentPlayer = otile
  }

  return nil
}

// View print current state of the game
func (g *Game) View() {
  render.Render()
  g.Board.view()
  if g.gameOver {
    fmt.Println("game over!")

    if g.winner != nil {
      fmt.Printf("player %v wins\n", g.winner)
    } else {
      fmt.Println("draw :(")
    }

  } else {
    fmt.Printf("player %v's turn\n", g.currentPlayer)
  }


  fmt.Println("____________________")
}

// Reset start a new game
func (g *Game) Reset() {
  g.Board.reset()
  g.gameOver = false
  g.currentPlayer = otile
  g.winner = nil

  for _, p := range g.players {
    p.reset()
  }
}

// GameOver is the game over?
func (g *Game) GameOver() bool {
  return g.gameOver
}
