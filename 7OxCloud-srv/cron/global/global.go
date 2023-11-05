package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/palp1tate/7OxCloud/7OxCloud-srv/cron/config"
	"github.com/palp1tate/7OxCloud/proto/transport"
)

var (
	Debug                  bool
	ServerConfig           config.ServerConfig
	NacosConfig            config.NacosConfig
	RedisClient            *redis.Client
	TransportServiceClient transportProto.TransportServiceClient
)
