package router

import (
	"github.com/gin-gonic/gin"
	_ "hangmango-web-api/model"
	"net/http"
)

func InitUserRouters(userGroup *gin.RouterGroup) {
	userGroup.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"total_count": 0,
			"data":        []int{},
		})
	})
}
