package controller

import (
	"github.com/gin-gonic/gin"
	db "hangmango-web-api/model"
	"net/http"
)

func StartNewGame(c *gin.Context) {
	userId, _ := c.Get("UserId")
	wordStr, err := db.StartNewGame(userId.(uint))
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"word": wordStr,
	})
	return
}
