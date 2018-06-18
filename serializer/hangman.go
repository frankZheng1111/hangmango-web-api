package serializer

import (
	db "hangmango-web-api/model"
)

type GuessingHangman struct {
	Id   int64  `json:"id"`
	Word string `json:"word"`
	Hp   int    `json:"hp"`
}

func SerializeGuessingHangman(hangman *db.Hangman) *GuessingHangman {
	guessingHangman := new(GuessingHangman)
	guessingHangman.Id = hangman.Id
	guessingHangman.Hp = hangman.Hp
	guessingHangman.Word = hangman.GameStr()
	return guessingHangman
}
