package api

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func sessionID(c *gin.Context) {
	sessionID, err := strconv.Atoi(c.Param("sessionID"))
	if err != nil {
		c.JSON(400, gin.H{
			"message":   "Invalid session id",
			"sessionID": sessionID,
		})
		return
	}

	g, ok := session[sessionID]
	if !ok {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("Invalid Session ID: %v", sessionID),
		})
		return
	}

	c.Set("game", g)
	c.Next()
}
