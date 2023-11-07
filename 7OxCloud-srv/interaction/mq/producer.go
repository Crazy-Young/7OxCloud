package mq

import (
	"github.com/palp1tate/7OxCloud/7OxCloud-srv/interaction/global"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

func SendMessage2MQ(body []byte, queueName string) (err error) {
	ch, err := global.RabbitMQ.Channel()
	if err != nil {
		zap.S().Warnf("failed to open a channel: %s", err.Error())
		return
	}
	q, _ := ch.QueueDeclare(queueName, true, false, false, false, nil)
	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         body,
	})
	if err != nil {
		zap.S().Warnf("failed to publish a message: %s", err.Error())
		return
	}
	return
}
