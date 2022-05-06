package main

import (
	"fmt"
	"math/rand"
	"tictactoe/game"
	"tictactoe/gym"
	"tictactoe/player"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func chooseStartingPlayer() int {
	if rand.Float64() > 0.5 {
		return 1
	}

	return 0
}

func main() {
	playerA := player.NewPlayer("player a", 0.3)
	playerB := player.NewPlayer("player b", 0.3)
	players := []*player.Player{playerA, playerB}
	scores := []float64{0, 0}

	env, err := gym.NewEnv("http://localhost:8080")
	if err != nil {
		panic(fmt.Sprintf("could not create a new env: %v", err.Error()))
	}

	for i := 0; i < 1000000; i++ {
		done := env.Done()
		playerIdx := chooseStartingPlayer()

		// Set player tiles
		players[playerIdx].SetTile(game.OTile)
		players[(playerIdx+1)%2].SetTile(game.XTile)
		obs := env.Observation()

		for !done {
			currentPlayer := players[playerIdx]
			actions := env.ActionSpace()

			action := currentPlayer.ChooseAction(actions, &obs)
			obs, done = env.Step(action)

			currentPlayer.AddState(obs)

			// next player's turn
			playerIdx = (playerIdx + 1) % 2
		}

		for i, p := range players {
			reward := env.Reward(p.Tile())
			p.FeedReward(reward)
			scores[i] += reward
		}

		fmt.Println(obs)

		obs = env.Reset()
	}
}
