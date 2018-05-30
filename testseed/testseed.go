package testseed

import (
	"github.com/jinzhu/gorm"
	"hangmango-web-api/config"
)

// 初始化数据
//
func InitTestDB(db *gorm.DB) {
	if config.Config.ENV != "test" {
		panic("invali env")
	}
	db.Exec("delete from users;")
}
