package main

import (
  "fmt"
  "math/rand"
  "time"

  "tictactoe/game"
)


func main() {
  rand.Seed(time.Now().Unix())

  g := game.NewGame()
  fmt.Printf("%v \n", g.Board.Moves())

  g.View()
  for ! g.GameOver()  {
    moves := g.Board.Moves()
    move := moves[rand.Intn(len(moves))]

    fmt.Printf("%v", move)

    g.Play(move.X(), move.Y())
    g.View()
  }
  g.View()

}
