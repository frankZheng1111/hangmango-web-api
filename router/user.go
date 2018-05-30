package router

import (
	"github.com/gin-gonic/gin"
	db "hangmango-web-api/model"
	"net/http"
)

func InitUserRouters(userGroup *gin.RouterGroup) {
	userGroup.GET("/", func(c *gin.Context) {
		db.CreateUser("test", "pass")
		c.JSON(http.StatusOK, gin.H{
			"total_count": 0,
			"data":        []int{},
		})
	})
}
