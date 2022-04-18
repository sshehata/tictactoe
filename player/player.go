package player

import (
	"math"
	"math/rand"
	"tictactoe/game"
	"tictactoe/gym"
)

// Player RL agent playing tictactoe
type Player struct {
	name        string
	states      []game.Position
	lr          float32
	expRate     float32 // exploration rate
	decayGamma  float32
	statesValue map[string]float64
	symbol      game.Tile
}

// NewPlayer create new player
func NewPlayer(name string, expRate float32) *Player {
	return &Player{
		name:        name,
		states:      make([]game.Position, 0, 9),
		lr:          0.2,
		expRate:     expRate,
		decayGamma:  0.9,
		statesValue: make(map[string]float64),
		symbol:      game.Undefined,
	}
}

// ChooseAction choose optimal actions from a set of
// possible moves or make an exploratory move
func (p *Player) ChooseAction(positions []game.Position, board gym.Board) game.Position {
	if rand.Float32() <= p.expRate {
		idx := rand.Intn(len(positions))
		return positions[idx]
	}

	max := math.Inf(-1)
	var action game.Position
	for _, pos := range positions {
		boardCopy := board
		boardCopy.Play(p.symbol, pos)
		value := p.statesValue[boardCopy.Hash()]
		if value > max {
			max = value
			action = pos
		}
	}
	return action
}

// SetTile assign the player it's symbol
func (p *Player) SetTile(t game.Tile) {
	p.symbol = t
}
