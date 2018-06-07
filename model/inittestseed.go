package model

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite" // just for test
	"hangmango-web-api/config"
	"log"
)

// 初始化数据(仅测试环境)
//
func InitTestDB() {
	if config.Config.ENV != "test" {
		panic("invalid env")
	}
	// Migrate the schema
	DB.AutoMigrate(&User{}, &Hangman{}, &HangmanGuessedLetter{})
	log.Println("AutoMigrate test db")

	DB.Delete(&User{})
	DB.Delete(&Hangman{})
	DB.Delete(&HangmanGuessedLetter{})

	DB.Create(&User{Id: 1, LoginName: "test-user-name", PasswordHash: "passwordhash"})
	DB.Create(&Hangman{Id: 1, UserId: 1, Word: "abandon"})
	DB.Create(&HangmanGuessedLetter{Id: 1, Letter: "a", HangmanId: 1})
}