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
	states      []gym.Board
	lr          float64
	expRate     float64 // exploration rate
	decayGamma  float64
	statesValue map[gym.Board]float64
	symbol      game.Tile
}

// NewPlayer create new player
func NewPlayer(name string, expRate float64) *Player {
	return &Player{
		name:        name,
		states:      make([]gym.Board, 0, 9),
		lr:          0.2,
		expRate:     expRate,
		decayGamma:  0.9,
		statesValue: make(map[gym.Board]float64),
		symbol:      game.Undefined,
	}
}

// ChooseAction choose optimal actions from a set of
// possible moves or make an exploratory move
func (p *Player) ChooseAction(positions []game.Position, board *gym.Board) game.Position {
	if rand.Float64() <= p.expRate {
		idx := rand.Intn(len(positions))
		return positions[idx]
	}

	max := math.Inf(-1)
	var action game.Position
	for _, pos := range positions {
		boardCopy := board
		boardCopy.Play(p.symbol, pos)
		value := p.statesValue[*boardCopy]
		if value > max {
			max = value
			action = pos
		}
	}
	return action
}

// FeedReward update policy based on reward
func (p *Player) FeedReward(reward float64) {
	for i := len(p.states) - 1; i >= 0; i-- {
		state := p.states[i]
		p.statesValue[state] += p.lr * (p.decayGamma*reward - p.statesValue[state])

	}
}

// SetTile assign the player it's symbol
func (p *Player) SetTile(t game.Tile) {
	p.symbol = t
}

// Tile get player symbol
func (p *Player) Tile() game.Tile {
	return p.symbol
}

// AddState add state to list of states player observes
func (p *Player) AddState(b gym.Board) {
	p.states = append(p.states, b)
}
