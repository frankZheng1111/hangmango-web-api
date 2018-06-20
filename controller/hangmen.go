package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hangmango-web-api/lib"
	db "hangmango-web-api/model"
	"hangmango-web-api/serializer"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type GuessLetter struct {
	Letter string `json:"letter" binding:"required"`
}

func StartNewGame(c *gin.Context) {
	userId, _ := c.Get("UserId")
	hangman := db.StartNewGame(userId.(int64))
	c.JSON(http.StatusOK, serializer.SerializeGuessingHangman(hangman))
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

	guessLetterKey := fmt.Sprintf("user-%v-guess-hangman-%v", userId, hangmanId)
	expiredIn := time.Duration(5) * time.Second
	timestamp, err := lib.GetRedisLock(guessLetterKey, expiredIn)
	if err != nil {
		panic(err)
	}
	if timestamp == 0 {
		MissingLockErrorResponse(c)
		return
	}
	_, err = hangman.Guess(guessLetter.Letter)
	lib.UnlockRedisLock(guessLetterKey, timestamp)
	if hangman.Status != "PLAYING" {
		user.UpdateScore(hangman.IsWin())
	}
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, serializer.SerializeGuessingHangman(hangman))
	return
}
