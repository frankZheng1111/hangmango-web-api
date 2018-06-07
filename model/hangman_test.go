package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStartNewGame(t *testing.T) {
	InitTestDB()
	hangman, _ := StartNewGame(1)
	assert.Equal(t, "abandon", hangman.Word)
	assert.Equal(t, "PLAYING", hangman.Status)
	assert.Equal(t, uint(1), hangman.UserId)
}

func TestAssociatedHangmenGuessedLetters(t *testing.T) {
	InitTestDB()
	hangman := new(Hangman)
	hangman.Id = 1
	DB.Where(hangman).Find(hangman)
	letters, _ := hangman.AssociatedHangmenGuessedLetters()
	assert.Equal(t, 1, len(letters))
	assert.Equal(t, "a", letters[0].Letter)

	hangman = new(Hangman)
	hangman.Id = 2
	DB.Where(hangman).Find(hangman)
	letters, err := hangman.AssociatedHangmenGuessedLetters()
	assert.Equal(t, 0, len(letters))
	assert.Nil(t, err)
}

func TestGameStr(t *testing.T) {
	InitTestDB()
	hangman := new(Hangman)
	hangman.Id = 1
	DB.Where(hangman).Find(hangman)
	gameStr, _ := hangman.GameStr()
	assert.Equal(t, "a*a****", gameStr)
}

func TestLeftHp(t *testing.T) {
	InitTestDB()
	hangman := new(Hangman)
	hangman.Id = 1
	DB.Where(hangman).Find(hangman)
	leftHp, _ := hangman.LeftHp()
	assert.Equal(t, 2, leftHp)
}
