package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/palp1tate/7OxCloud/7OxCloud-srv/interaction/config"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

var (
	DB           *gorm.DB
	Debug        bool
	RedisClient  *redis.Client
	ServerConfig config.ServerConfig
	NacosConfig  config.NacosConfig
	RabbitMQ     *amqp.Connection
)
