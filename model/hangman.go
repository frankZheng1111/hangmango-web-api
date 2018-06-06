package model

import (
	"hangmango-web-api/config"
	"math/rand"
	"regexp"
	"time"
)

type Hangman struct {
	Base
	Id                    uint   `gorm:"column:id; primary_key"`
	UserId                uint   `gorm:"column:user_id"`
	Word                  string `gorm:"column:word"`
	Status                string `gorm:"column:status;default:PLAYING"`
	HangmenGuessedLetters []HangmanGuessedLetter
}

func (hangman *Hangman) Guess(letter string) (hangmanGuessedLetter *HangmanGuessedLetter, err error) {
	result := DB.Create(&HangmanGuessedLetter{Letter: letter, HangmanId: hangman.Id})
	err = result.Error
	if err != nil {
		return
	}
	hangmanGuessedLetter = result.Value.(*HangmanGuessedLetter) // must be ptr
	return
}

func (hangman *Hangman) AssociatedHangmenGuessedLetters() ([]*HangmanGuessedLetter, error) {
	hangmanGuessedLetters := []*HangmanGuessedLetter{}
	err := DB.Model(hangman).Association("HangmenGuessedLetters").Find(&hangmanGuessedLetters).Error
	if err != nil {
		return nil, err
	}
	return hangmanGuessedLetters, nil
}

func StartNewGame(userId uint) (gameStr string, err error) {
	source := rand.NewSource(time.Now().Unix())
	randMachine := rand.New(source)
	randIndex := randMachine.Intn(len(config.Config.Hangman.Dictionary) - 1)
	word := config.Config.Hangman.Dictionary[randIndex]
	err = DB.Create(&Hangman{UserId: userId, Word: word}).Error
	if err != nil {
		return
	}
	var re = regexp.MustCompile(`[a-zA-Z]`)
	gameStr = re.ReplaceAllString(word, `*`)
	return
}
