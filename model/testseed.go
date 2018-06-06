package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // just for test
	"hangmango-web-api/config"
	"log"
)

// 初始化数据(仅测试环境)
//
func InitTestDB(db *gorm.DB) {
	if config.Config.ENV != "test" {
		panic("invali env")
	}
	// Migrate the schema
	db.AutoMigrate(&User{}, &Hangman{}, &HangmanGuessedLetter{})
	log.Println("AutoMigrate test db")

	DB.Delete(&User{})
	DB.Delete(&Hangman{})
	DB.Delete(&HangmanGuessedLetter{})

}
