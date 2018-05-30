package router

import (
	"github.com/gin-gonic/gin"
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
