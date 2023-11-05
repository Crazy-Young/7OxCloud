package initialize

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/palp1tate/7OxCloud/7OxCloud-srv/video/global"
)

func InitRedis() {
	host := global.ServerConfig.Redis.Host
	port := global.ServerConfig.Redis.Port
	global.RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: "",
		DB:       1, // 存储用户浏览记录
	})
}
