package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // just for test
	"hangmango-web-api/config"
	"log"
	"time"
)

type Base struct {
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

type Paginate struct {
	Page     int
	PageSize int
}

func (paginate *Paginate) ParseToLimitAndOffset() (limit int, offset int) {
	if 0 >= paginate.Page || 0 >= paginate.PageSize {
		paginate.Page = 1
		paginate.PageSize = 30
	}
	limit = paginate.PageSize
	offset = (paginate.Page - 1) * paginate.PageSize
	return
}

var DB *gorm.DB
var err error

func init() {
	InitModel()
}

func InitModel() {
	DB, err = gorm.Open(config.Config.GORM.Driver, config.Config.GORM.Open)
	if err != nil {
		panic("failed to connect database")
	}

	DB.DB().SetMaxIdleConns(config.Config.GORM.MaxIdle)
	DB.DB().SetMaxOpenConns(config.Config.GORM.MaxOpen)

	if config.Config.ENV == "test" {
		// Migrate the schema
		DB.AutoMigrate(&User{})
		log.Println("AutoMigrate test db")
	}

	log.Println("Init Model Complete")
}
