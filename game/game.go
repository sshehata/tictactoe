package game

import (
	"errors"
	"fmt"
	"strings"
)

// Game a game of tictactoe
type Game struct {
	board         *Board
	currentPlayer Tile
	moveCount     int
	gameOver      bool
	winner        *player
	players       [3]*player
}

// NewGame Start a new tictactoe game
func NewGame() *Game {
	game := Game{
		board: newBoard(),
		players: [3]*player{
			newPlayer(Undefined),
			newPlayer(OTile),
			newPlayer(XTile),
		},
	}
	game.Reset()
	return &game
}

// Play current player makes a move
func (g *Game) Play(x, y int) error {
	if g.gameOver {
		return errors.New("the game is over")
	}

	p := g.players[g.currentPlayer]
	err := g.board.play(g.players[g.currentPlayer].tile, Position{x, y})
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
	case OTile:
		g.currentPlayer = XTile
	case XTile:
		g.currentPlayer = OTile
	}

	return nil
}

// View print current state of the game
func (g *Game) String() string {
	var writer strings.Builder
	fmt.Fprintf(&writer, "%v\n", g.board.String())
	fmt.Fprintf(&writer, "____________________")
	return writer.String()
}

// Reset start a new game
func (g *Game) Reset() {
	g.board.reset()
	g.gameOver = false
	g.currentPlayer = OTile
	g.winner = g.players[0]

	for _, p := range g.players {
		p.reset()
	}
}

// GameOver is the game over?
func (g *Game) GameOver() bool {
	return g.gameOver
}

// CurrentPlayer get the player playing this turn
func (g *Game) CurrentPlayer() Tile {
	return g.currentPlayer
}

// Winner get winner
func (g *Game) Winner() *player {
	return g.winner
}

// Board return the current game board state
func (g *Game) Board() *Board {
	return g.board
}

// IsDraw return true if the game is draw
func (g *Game) IsDraw() bool {
	return g.gameOver && g.winner == g.players[0]
}
