package router

import (
	"github.com/gin-gonic/gin"
	"hangmango-web-api/controller"
	"net/http"
)

func InitUserRouters(userGroup *gin.RouterGroup) {
	userGroup.GET("/probe", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "success"})
	})
	userGroup.POST("/", CommonPanicHandle(controller.UserSignUp))
	userGroup.POST("/signin", CommonPanicHandle(controller.UserSignIn))
	userGroup.Use(controller.CommonPanicHandle(controller.ValidAuthToken))
	userGroup.GET("/best-users", CommonPanicHandle(controller.GetBestUsers))
}
