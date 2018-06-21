package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStartNewGame(t *testing.T) {
	InitTestDB()
	hangman := StartNewGame(1)
	assert.Equal(t, "abandon", hangman.Word)
	assert.Equal(t, "PLAYING", hangman.Status)
	assert.Equal(t, int64(1), hangman.UserId)
}

func TestAssociatedHangmanGuessedLetters(t *testing.T) {
	InitTestDB()
	hangman := new(Hangman)
	hangman.Id = 1
	DB.Where(hangman).Find(hangman)
	letters := hangman.AssociatedHangmanGuessedLetters()
	assert.Equal(t, 1, len(letters))
	assert.Equal(t, "a", letters[0].Letter)

	hangman = new(Hangman)
	hangman.Id = 2
	DB.Where(hangman).Find(hangman)
	letters = hangman.AssociatedHangmanGuessedLetters()
	assert.Equal(t, 0, len(letters))
	assert.Nil(t, err)
}

func TestGameStr(t *testing.T) {
	InitTestDB()
	hangman := new(Hangman)
	hangman.Id = 1
	DB.Where(hangman).Find(hangman)
	gameStr := hangman.GameStr()
	assert.Equal(t, "a*a****", gameStr)
}

func TestIsAlive(t *testing.T) {
	InitTestDB()
	hangman := new(Hangman)
	hangman.Id = 2
	DB.Where(hangman).Find(hangman)
	assert.True(t, hangman.IsAlive())
}

func TestIsWin(t *testing.T) {
	InitTestDB()
	hangman := new(Hangman)
	hangman.Id = 2
	DB.Where(hangman).Find(hangman)
	assert.False(t, hangman.IsWin())
	hangman = new(Hangman)
	hangman.Id = 3
	DB.Where(hangman).Find(hangman)
	assert.True(t, hangman.IsWin())
}

func TestIsGuessFail(t *testing.T) {
	InitTestDB()
	hangman := new(Hangman)
	hangman.Id = 3
	DB.Where(hangman).Find(hangman)
	_, err := hangman.Guess("a")
	assert.Equal(t, "AlreadyFinish", err.Error())
	//
	hangman = new(Hangman)
	hangman.Id = 4
	DB.Where(hangman).Find(hangman)
	_, err = hangman.Guess("esss")
	assert.Equal(t, "InvalidLetter", err.Error())
	_, err = hangman.Guess("c")
	assert.Nil(t, err)
	_, err = hangman.Guess("c")
	assert.Nil(t, err)
	_, err = hangman.Guess("d")
	assert.Equal(t, "AlreadyFinish", err.Error())
}
