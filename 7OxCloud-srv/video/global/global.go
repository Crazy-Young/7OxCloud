package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/palp1tate/7OxCloud/7OxCloud-srv/video/config"
	"github.com/palp1tate/7OxCloud/proto/transport"
	"gorm.io/gorm"
)

var (
	DB                     *gorm.DB
	Debug                  bool
	ServerConfig           config.ServerConfig
	NacosConfig            config.NacosConfig
	RedisClient            *redis.Client
	TransportServiceClient transportProto.TransportServiceClient
)
