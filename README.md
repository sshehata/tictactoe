# tictactoe
Tic Tac Toe is a popular children's game with simple rules.

![](https://upload.wikimedia.org/wikipedia/commons/7/7d/Tic-tac-toe-animated.gif)

This repository investigates the use of different reinforcement learning techniques to train an agent that is proficient at wining Tic Tac Toe.

# Usage

### TicTacToe Server
Run the following to start a tictactoe server on port `8080`"
```
go run ./cmd/tictactoe/main.go
```

### Train

Run the following to start a training session:
```
go run ./cmd/train/main.go
```
This will create two new files:
* _playera.bin_
* _playerb.bin_

trained by playing the two agents against each other.

# Additional Info

### Training

Training runs for 1000 epochs each with 1000 games. The player policy is persisted to disk after each epoch.

### Packages

* __game__ Tic Tac Toe game engine.
* __api__ RESTful api to interface with the game engine.
* __player__ RL agent logic (Q Agent).
* __gym__ OpenAI inspired gym environment to support training of agents against the game engine.
* __cmd__ Entry points for the different packages (training, game server).

# Roadmap

* [ ] Web-based GUI for human vs agent games.
* [ ] Support more RL algorithms.
* [ ] Usability improvements (e.g. Command-line arguments for hard-coded parameters, better logging, etc...)

# Acknowledgments

* Heavely inspired by [Reinforcement Learning — Implement TicTacToe](https://towardsdatascience.com/reinforcement-learning-implement-tictactoe-189582bea542)
