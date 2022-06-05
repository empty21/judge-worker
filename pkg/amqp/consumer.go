package amqp

import (
	"bytes"
	"encoding/json"
	"github.com/streadway/amqp"
	"judger/pkg/config"
	"judger/pkg/logger"
)

var consumerChannel *amqp.Channel

func initConsumer() {
	var messages <-chan amqp.Delivery

	connection, err := amqp.Dial(config.Config.AMQPUri)
	if err == nil {
		consumerChannel, err = connection.Channel()
	}
	if err == nil {
		_, err = consumerChannel.QueueDeclare(
			JudgerServiceQueue,
			true,
			false,
			false,
			false,
			nil,
		)
	}
	if err == nil {
		messages, err = consumerChannel.Consume(
			JudgerServiceQueue,
			"",
			false,
			false,
			false,
			false,
			nil,
		)
	}

	if err != nil {
		panic(err)
	}
	logger.Logger.Info("Successfully connected to RabbitMQ")
	logger.Logger.Info("Waiting for messages")
	forever := make(chan bool)
	go func() {
		for message := range messages {
			messageCompact := new(bytes.Buffer)
			if err = json.Compact(messageCompact, message.Body); err != nil {
				logger.AMQPLogger.Error(err)
				logger.AMQPLogger.Info("CONSUMER: ", message.Body)
			} else {
				logger.AMQPLogger.Info("CONSUMER: ", messageCompact)
			}

			err = handleMessage(message)
			if err != nil {
				_ = message.Nack(false, false)
			} else {
				_ = message.Ack(false)
			}
		}
	}()
	<-forever
}
