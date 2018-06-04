package model

import (
	"hangmango-web-api/config"
	"math/rand"
	"regexp"
	"time"
)

type Hangman struct {
	Base
	Id     uint   `gorm:"column:id; primary_key"`
	UserId uint   `gorm:"column:user_id"`
	Word   string `gorm:"column:word"`
	Status string `gorm:"column:status;default:PLAYING"`
}

func StartNewGame(userId uint) (gameStr string, err error) {
	source := rand.NewSource(time.Now().Unix())
	randMachine := rand.New(source)
	randIndex := randMachine.Intn(len(config.Config.Dictionary) - 1)
	word := config.Config.Dictionary[randIndex]
	err = DB.Create(&Hangman{UserId: userId, Word: word}).Error
	if err != nil {
		return
	}
	var re = regexp.MustCompile(`[a-zA-Z]`)
	gameStr = re.ReplaceAllString(word, `*`)
	return
}
