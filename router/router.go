package router

import (
	"github.com/gin-gonic/gin"
	"hangmango-web-api/controller"
	"net/http"
)

var CommonPanicHandle func(action func(c *gin.Context)) func(c *gin.Context) = controller.CommonPanicHandle

func InitRouters(version *gin.RouterGroup) {
	version.GET("/probe", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "success"})
	})
	InitUserRouters(version.Group("/users"))
	InitHangmanRouters(version.Group("/hangmen"))
}
