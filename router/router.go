package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouters(version *gin.RouterGroup) {
	version.GET("/probe", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "success",
		})
	})
}
