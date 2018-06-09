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
	DB.Create(&Hangman{Id: 1, UserId: 1, Word: "abandon", Hp: 2})
	DB.Create(&Hangman{Id: 2, UserId: 1, Word: "abandon", Hp: 2})
	DB.Create(&Hangman{Id: 3, UserId: 1, Word: "a", Hp: 2, Status: "WIN"})
	DB.Create(&Hangman{Id: 4, UserId: 1, Word: "ant", Hp: 2})
	DB.Create(&HangmanGuessedLetter{Id: 1, Letter: "a", HangmanId: 1})
	DB.Create(&HangmanGuessedLetter{Id: 2, Letter: "a", HangmanId: 3})
}
