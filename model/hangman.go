package model

import (
	"hangmango-web-api/config"
	"math/rand"
	"strings"
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

func (hangman *Hangman) GuessedLettersMap() (lettersMap map[string]bool) {
	var hangmanGuessedLetters []*HangmanGuessedLetter
	lettersMap = make(map[string]bool)
	hangmanGuessedLetters, err = hangman.AssociatedHangmenGuessedLetters()
	if err != nil {
		return
	}
	for _, hangmanGuessedLetter := range hangmanGuessedLetters {
		lettersMap[hangmanGuessedLetter.Letter] = strings.Contains(hangman.Word, hangmanGuessedLetter.Letter)
	}
	return
}

func (hangman *Hangman) GameStr() (gameStr string, err error) {
	guessedLetters := hangman.GuessedLettersMap()
	for _, wordLetterRune := range hangman.Word {
		wordLetter := string(wordLetterRune)
		if _, ok := guessedLetters[wordLetter]; ok {
			gameStr += wordLetter
		} else {
			gameStr += "*"
		}
	}
	return
}

func StartNewGame(userId uint) (hangman *Hangman, err error) {
	source := rand.NewSource(time.Now().Unix())
	randMachine := rand.New(source)
	randIndex := randMachine.Intn(len(config.Config.Hangman.Dictionary) - 1)
	word := config.Config.Hangman.Dictionary[randIndex]
	result := DB.Create(&Hangman{UserId: userId, Word: word})
	err = result.Error
	if err != nil {
		return
	}
	hangman = result.Value.(*Hangman)
	return
}
