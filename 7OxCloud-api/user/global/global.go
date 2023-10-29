package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-redis/redis/v8"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/user/config"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/user/proto"
)

var (
	Debug      bool
	Translator ut.Translator

	RedisClient *redis.Client

	ServerConfig *config.ServerConfig

	NacosConfig *config.NacosConfig

	UserServiceClient proto.UserServiceClient
)
