package serializer

import (
	"github.com/stretchr/testify/assert"
	db "hangmango-web-api/model"
	"testing"
)

func TestSerializeGuessingHangman(t *testing.T) {
	hangman := new(db.Hangman)
	hangman.Id = 1
	hangman.Word = "test"
	hangman.Hp = 1

	guessingHangman := SerializeGuessingHangman(hangman)
	assert.Equal(t, int64(1), guessingHangman.Id)
	assert.Equal(t, 1, guessingHangman.Hp)
	assert.Equal(t, "****", guessingHangman.Word)
}
