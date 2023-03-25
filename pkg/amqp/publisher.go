package amqp

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"judger/pkg/config"
	"judger/pkg/logger"
)

var publisherChannel *amqp.Channel

func initPublisher() {
	connection, err := amqp.Dial(config.Config.AMQPUri)
	if err == nil {
		publisherChannel, err = connection.Channel()
	}
	if err == nil {
		_, err = publisherChannel.QueueDeclare(
			config.Config.ResultQueue,
			true,
			false,
			false,
			false,
			nil,
		)
	}
	if err == nil {
		err = publisherChannel.ExchangeDeclare(
			"common",
			"direct",
			true,
			false,
			false,
			false,
			nil,
		)
	}
	if err == nil {
		err = publisherChannel.QueueBind(config.Config.ResultQueue, JudgeTaskUpdate, "common", false, nil)
	}
	if err != nil {
		panic(err)
	}
}

func Publish(key Pattern, message Message) error {
	msg, err := json.Marshal(message)
	if err == nil {
		err = publisherChannel.Publish("common", string(key), false, false, amqp.Publishing{
			Headers:     nil,
			ContentType: "application/json",
			Body:        msg,
		})
	}
	if err != nil {
		return err
	}
	logger.AMQPLogger.Info("PUBLISHER: ", string(msg))
	return nil
}
