package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hangmango-web-api/config"
	db "hangmango-web-api/model"
	"hangmango-web-api/router"
)

func main() {
	if config.Config.ENV == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	defer db.DB.Close()
	ginRouter := gin.Default()

	v1 := ginRouter.Group("/v1")
	router.InitRouters(v1)

	ginRouter.Run(fmt.Sprintf(":%d", config.Config.Server.Port))
}
