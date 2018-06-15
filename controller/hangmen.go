package controller

import (
	"github.com/gin-gonic/gin"
	db "hangmango-web-api/model"
	"net/http"
	"strconv"
	"strings"
)

type GuessLetter struct {
	Letter string `json:"letter" binding:"required"`
}

func StartNewGame(c *gin.Context) {
	userId, _ := c.Get("UserId")
	hangman := db.StartNewGame(userId.(int64))
	gameStr := hangman.GameStr()
	c.JSON(http.StatusOK, gin.H{
		"id":   hangman.Id,
		"hp":   hangman.Hp,
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
	guessLetter := new(GuessLetter)
	if err = c.BindJSON(guessLetter); err != nil {
		if !strings.Contains(err.Error(), "validation") {
			panic(err)
		}
		ValidationErrorResponse(c)
		return
	}

	userId, _ := c.Get("UserId")
	user, err := db.GetUserById(userId.(int64))
	if err != nil {
		panic(err)
	}
	hangman, err := user.HangmenById(int64(hangmanId))
	if err != nil {
		panic(err)
	}

	_, err = hangman.Guess(guessLetter.Letter)
	if err != nil {
		panic(err)
	}
	gameStr := hangman.GameStr()
	c.JSON(http.StatusOK, gin.H{
		"id":   hangman.Id,
		"hp":   hangman.Hp,
		"word": gameStr,
	})
	return
}
