package api

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	"tictactoe/game"
)

var session map[int]*game.Game

func init() {
	session = make(map[int]*game.Game)
}

func newGame(c *gin.Context) {
	sessionID := len(session) + 1
	game := game.NewGame()
	session[sessionID] = game
	c.JSON(200, gin.H{
		"sessionID": sessionID,
		"turn":      game.CurrentPlayer().String(),
		"board":     game.String(),
	})
}

func play(c *gin.Context) {
	sessionID, err := strconv.Atoi(c.Param("sessionID"))
	if err != nil {
		c.JSON(400, gin.H{
			"message":   "Invalid session id",
			"sessionID": sessionID,
		})
		return
	}

	g, ok := session[sessionID]
	fmt.Println(ok)
	if !ok {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("Invalid Session ID: %v", sessionID),
		})
		return
	}

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
		"board":  g.String(),
		"player": g.CurrentPlayer().String(),
		"status": g.GameOver(),
		"winner": g.Winner().String(),
	})
}

func reset(c *gin.Context) {
	sessionID, err := strconv.Atoi(c.Param("sessionID"))
	if err != nil {
		c.JSON(400, gin.H{
			"message":   "Invalid session id",
			"sessionID": sessionID,
		})
		return
	}

	game, ok := session[sessionID]
	fmt.Println(ok)
	if !ok {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("Invalid Session ID: %v", sessionID),
		})
		return
	}

	game.Reset()
}

// Listen start listener server
func Listen() {
	r := gin.Default()
	r.GET("/", newGame)
	r.POST("/:sessionID", play)
	r.POST("/:sessionID/reset", reset)

	r.Run()
}
