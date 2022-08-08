package bootstrap

import (
	"fmt"
	"go-simple/pkg/config"
	"go-simple/pkg/redis"
)

func SetupRedis() {
	redis.ConnectRedis(fmt.Sprintf("%v:%v", config.GetString("redis.host"),
		config.GetString("redis.port")),
		config.GetString("redis.username"),
		config.GetString("redis.password"),
		config.GetInt("redis.password"),
	)
}