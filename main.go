package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
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

	store, err := redis.NewStore(10, config.Config.Redis.Network, config.Config.Redis.Address, config.Config.Redis.Password, []byte(config.Config.Redis.AuthKey))
	if err != nil {
		panic(err)
	}

	ginRouter.Use(sessions.Sessions("hangmango-web", store))

	v1 := ginRouter.Group("/v1")
	router.InitRouters(v1)

	ginRouter.Run(fmt.Sprintf(":%d", config.Config.Server.Port))
}
