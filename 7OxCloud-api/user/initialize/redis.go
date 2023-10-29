package initialize

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/user/global"
)

func InitRedis() {
	host := global.ServerConfig.Redis.Host
	port := global.ServerConfig.Redis.Port
	global.RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port), // Redis 服务器地址
		Password: "",                               // Redis 访问密码（如果有的话）
		DB:       0,                                // Redis 数据库索引
	})
}
