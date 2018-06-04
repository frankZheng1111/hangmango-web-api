package router

import (
	"github.com/gin-gonic/gin"
	"hangmango-web-api/controller"
	"net/http"
)

func InitHangmanRouters(userGroup *gin.RouterGroup) {
	userGroup.GET("/probe", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "success"})
	})
	userGroup.Use(CommonPanicHandle(controller.ValidAuthToken))
	userGroup.POST("/", CommonPanicHandle(controller.StartNewGame))
}
