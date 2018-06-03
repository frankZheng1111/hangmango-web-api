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
	userGroup.POST("/", controller.UserSignUp)
	userGroup.POST("/signin", controller.UserSignIn)
	userGroup.Use(controller.ValidAuthToken)
	userGroup.GET("/best-users", controller.GetBestUsers)
}
