package model

import (
	"errors"
	"hangmango-web-api/config"
	"hangmango-web-api/lib"
	"math/rand"
	"strings"
	"time"
)

type Hangman struct {
	Base
	Id                    int64  `gorm:"column:id; primary_key"`
	UserId                int64  `gorm:"column:user_id"`
	Hp                    int    `gorm:"column:hp"`
	Word                  string `gorm:"column:word"`
	Status                string `gorm:"column:status;default:PLAYING"`
	HangmenGuessedLetters []HangmanGuessedLetter
}

var HangmanSnowflake *lib.Snowflake

func init() {
	HangmanSnowflake = lib.NewSnowflake()
}

func (hangman *Hangman) Guess(letter string) (hangmanGuessedLetter *HangmanGuessedLetter, err error) {
	if len(letter) != 1 {
		return nil, errors.New("InvalidLetter")
	}
	if hangman.Status != "PLAYING" {
		return nil, errors.New("AlreadyFinish")
	}
	tx := DB.Begin()
	result := tx.Create(&HangmanGuessedLetter{
		Id:        HangmanGuessedLetterSnowflake.Id(),
		Letter:    letter,
		HangmanId: hangman.Id,
	})
	err = result.Error
	if err != nil {
		tx.Rollback()
		return
	}
	if !strings.Contains(hangman.Word, letter) || hangman.GuessedLettersMap()[letter] > 1 {
		hangman.Hp--
	}
	if !hangman.IsAlive() {
		hangman.Status = "FAIL"
	}
	if hangman.IsWin() {
		hangman.Status = "WIN"
	}
	err = tx.Save(hangman).Error
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	hangmanGuessedLetter = result.Value.(*HangmanGuessedLetter) // must be ptr
	return
}

func (hangman *Hangman) AssociatedHangmenGuessedLetters() []*HangmanGuessedLetter {
	hangmanGuessedLetters := []*HangmanGuessedLetter{}
	err := DB.Model(hangman).Association("HangmenGuessedLetters").Find(&hangmanGuessedLetters).Error
	if err != nil {
		panic(err)
	}
	return hangmanGuessedLetters
}

func (hangman *Hangman) GuessedLettersMap() (lettersMap map[string]int) {
	var hangmanGuessedLetters []*HangmanGuessedLetter
	lettersMap = make(map[string]int)
	hangmanGuessedLetters = hangman.AssociatedHangmenGuessedLetters()
	for _, hangmanGuessedLetter := range hangmanGuessedLetters {
		lettersMap[hangmanGuessedLetter.Letter]++
	}
	return
}

func (hangman *Hangman) IsWin() bool {
	return !strings.Contains(hangman.GameStr(), "*")
}

func (hangman *Hangman) IsAlive() bool {
	return hangman.Hp > 0
}

func (hangman *Hangman) GameStr() (gameStr string) {
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

func StartNewGame(userId int64) (hangman *Hangman) {
	source := rand.NewSource(time.Now().Unix())
	randMachine := rand.New(source)
	randIndex := randMachine.Intn(len(config.Config.Hangman.Dictionary) - 1)
	word := config.Config.Hangman.Dictionary[randIndex]
	hp := config.Config.Hangman.Hp
	result := DB.Create(&Hangman{
		Id:     HangmanSnowflake.Id(),
		UserId: userId,
		Word:   word,
		Hp:     hp,
	})
	err = result.Error
	if err != nil {
		panic(err)
	}
	hangman = result.Value.(*Hangman)
	return
}
