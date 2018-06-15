package lib

import (
	"github.com/go-redis/redis"
	"hangmango-web-api/config"
)

var Client *redis.Client

func init() {
	InitRedis()
}

func InitRedis() {
	Client = redis.NewClient(&redis.Options{
		Network:  config.Config.Redis.Network,
		Addr:     config.Config.Redis.Address,
		Password: config.Config.Redis.Password,
		DB:       config.Config.Redis.DB,
	})
}
