package main

import (
	"fmt"
	"math/rand"
	"tictactoe/game"
	"tictactoe/gym"
	"tictactoe/player"
)

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

	env, err := gym.NewEnv("http://localhost:8080")
	if err != nil {
		panic(fmt.Sprintf("could not create a new env: %v", err.Error()))
	}

	for i := 0; i < 1000; i++ {
		gameOver := env.GameOver()
		currentPlayer := chooseStartingPlayer()

		players[currentPlayer].SetTile(game.OTile)
		players[(currentPlayer+1)%2].SetTile(game.XTile)
		obs := env.Observation()
		fmt.Println(obs)

		for !gameOver {
			moves := env.Moves()
			fmt.Println(moves)

			// next player's turn
			currentPlayer = (currentPlayer + 1) % 2
		}
	}
}
