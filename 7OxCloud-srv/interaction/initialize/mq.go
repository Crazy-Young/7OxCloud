package initialize

import (
	"fmt"

	"github.com/palp1tate/7OxCloud/7OxCloud-srv/interaction/global"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

func InitRabbitMQ() {
	connString := fmt.Sprintf("amqp://%s:%s@%s:%d/",
		global.ServerConfig.RabbitMQ.User,
		global.ServerConfig.RabbitMQ.Password,
		global.ServerConfig.RabbitMQ.Host,
		global.ServerConfig.RabbitMQ.Port,
	)
	conn, err := amqp.Dial(connString)
	if err != nil {
		zap.S().Panic("failed to connect rabbitmq:" + err.Error())
	}
	global.RabbitMQ = conn
}
