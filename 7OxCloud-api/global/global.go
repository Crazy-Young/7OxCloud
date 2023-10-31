package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-redis/redis/v8"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/config"
	"github.com/palp1tate/7OxCloud/proto/user"
	"github.com/palp1tate/7OxCloud/proto/video"
)

var (
	Debug      bool
	Translator ut.Translator

	RedisClient *redis.Client

	ServerConfig *config.ServerConfig

	NacosConfig *config.NacosConfig

	UserServiceClient  userProto.UserServiceClient
	VideoServiceClient videoProto.VideoServiceClient
)
