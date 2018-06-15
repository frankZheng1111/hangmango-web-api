package model

type HangmanGuessedLetter struct {
	Base
	Id        int64   `gorm:"column:id; primary_key"`
	HangmanId int64   `gorm:"column:hangman_id"`
	Letter    string `gorm:"column:letter"`
}
