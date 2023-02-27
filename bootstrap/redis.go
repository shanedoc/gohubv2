package bootstrap

import (
	"fmt"
	"gohubv2/pkg/config"
	"gohubv2/pkg/redis"
)

// SetupRedis 初始化 Redis
func SetUpRedis() {
	redis.ConnectRedis(
		fmt.Sprintf("%v:%v", config.GetString("redis.host"), config.GetString("redis.port")),
		config.GetString("redis.username"),
		config.GetString("redis.password"),
		config.GetInt("redis.database"),
	)
}
