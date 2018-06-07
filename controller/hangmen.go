package controller

import (
	"github.com/gin-gonic/gin"
	db "hangmango-web-api/model"
	"net/http"
	"strconv"
)

func StartNewGame(c *gin.Context) {
	userId, _ := c.Get("UserId")
	hangman, err := db.StartNewGame(userId.(uint))
	if err != nil {
		panic(err)
	}
	gameStr, err := hangman.GameStr()
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"word": gameStr,
	})
	return
}

func GuessALetter(c *gin.Context) {
	hangmanId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ValidationErrorResponse(c)
		return
	}

	user, err := db.GetUserById(3)
	if err != nil {
		panic(err)
	}
	hangman, err := user.HangmenById(uint(hangmanId))
	if err != nil {
		panic(err)
	}

	letter, err := hangman.Guess("z")
	c.JSON(http.StatusOK, gin.H{
		"msg": letter.Id,
	})
	return
}
