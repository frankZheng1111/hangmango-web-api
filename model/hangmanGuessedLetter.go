package model

type HangmanGuessedLetter struct {
	Base
	Id        uint   `gorm:"column:id; primary_key"`
	HangmanId uint   `gorm:"column:hangman_id"`
	Letter    string `gorm:"column:letter"`
}
