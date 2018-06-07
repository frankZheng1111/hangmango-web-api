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
