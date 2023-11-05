package mq

import (
	"context"

	"github.com/palp1tate/7OxCloud/7OxCloud-srv/interaction/global"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

func ConsumeMessage(ctx context.Context, queueName string) (msg <-chan amqp.Delivery, err error) {
	ch, err := global.RabbitMQ.Channel()
	if err != nil {
		zap.S().Warnf("failed to open a channel: %s", err.Error())
		return
	}
	q, _ := ch.QueueDeclare(queueName, true, false, false, false, nil)
	err = ch.Qos(1, 0, false)
	return ch.Consume(q.Name, "", false, false, false, false, nil)
}
