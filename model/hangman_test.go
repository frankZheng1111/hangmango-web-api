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
}
