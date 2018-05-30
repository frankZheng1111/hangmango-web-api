package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hangmango-web-api/config"
	"hangmango-web-api/router"
)

func main() {
	ginRouter := gin.Default()

	v1 := ginRouter.Group("/v1")
	router.InitRouters(v1)

	ginRouter.Run(fmt.Sprintf(":%d", config.Config.Server.Port))
}
