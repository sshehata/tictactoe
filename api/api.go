package api

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"tictactoe/game"
)

var session map[int]*game.Game

func init() {
	session = make(map[int]*game.Game)
}

func newGame(c *gin.Context) {
	sessionID := len(session) + 1
	g := game.NewGame()
	session[sessionID] = g
	c.JSON(200, gin.H{
		"sessionID": sessionID,
		"turn":      g.CurrentPlayer(),
		"board":     g.Board(),
	})
}

func play(c *gin.Context) {
	g := c.MustGet("game").(*game.Game)

	var p game.Position
	if err := c.BindJSON(&p); err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("Invalid move parameters: %v", err.Error()),
		})
		return
	}

	if err := g.Play(p.X, p.Y); err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("Invalid move: %v", err.Error()),
		})
		return
	}

	c.JSON(200, gin.H{
		"board":  g.Board(),
		"player": g.CurrentPlayer().String(),
		"status": g.GameOver(),
		"winner": g.Winner().String()})
}

func reset(c *gin.Context) {
	g := c.MustGet("game").(*game.Game)
	g.Reset()

	c.JSON(200, gin.H{
		"board":  g.Board(),
		"player": g.CurrentPlayer(),
		"status": g.GameOver(),
		"winner": g.Winner().Tile(),
	})
}

func moves(c *gin.Context) {
	g := c.MustGet("game").(*game.Game)
	g.Board().Moves()

	c.JSON(200, gin.H{
		"moves": g.Board().Moves()})
}

func gameOver(c *gin.Context) {
	g := c.MustGet("game").(*game.Game)
	c.JSON(200, gin.H{
		"gameOver": g.GameOver(),
	})
}

func winner(c *gin.Context) {
	g := c.MustGet("game").(*game.Game)

	if !g.GameOver() {
		c.JSON(400, gin.H{
			"error": "game is not over.",
		})
		return
	}

	if g.IsDraw() {
		c.JSON(200, gin.H{
			"winner": game.Undefined,
		})
	} else {
		c.JSON(200, gin.H{
			"winner": g.Winner().Tile(),
		})
	}
}

func state(c *gin.Context) {
	g := c.MustGet("game").(*game.Game)
	c.JSON(200, gin.H{
		"board": g.Board(),
	})
}

// Listen start listener server
func Listen() {
	r := gin.Default()
	r.GET("/", newGame)

	r.Use(sessionID)
	{
		r.GET("/:sessionID/moves/", moves)
		r.GET("/:sessionID/winner", winner)
		r.GET("/:sessionID/gameover", gameOver)
		r.GET("/:sessionID/", state)
		r.POST("/:sessionID/", play)
		r.POST("/:sessionID/reset", reset)
	}

	r.Run()
}
