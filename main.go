package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hangmango-web-api/config"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/probe", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "success",
		})
	})

	router.Run(fmt.Sprintf(":%d", config.Config.Server.Port))
}
