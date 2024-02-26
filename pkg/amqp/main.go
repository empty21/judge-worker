package amqp

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"judger/pkg/config"
	"judger/pkg/log"
	"os"
	"time"
)

var connection *amqp.Connection
var channel *amqp.Channel
var channelClosed chan *amqp.Error
var retryCount = 0

func Init() {
	c()
	for {
		select {
		case err := <-channelClosed:
			log.Error("AMQP connection lost (%v), reconnect after 10 seconds", err)
			<-time.After(10 * time.Second)
			c()
		}
	}
}

func c() {
	var err error
	connection, channel, err = connect(config.Config.AMQPUri)
	if err != nil {
		if retryCount < 5 {
			retryCount++
			log.Error("AMQP connection failed (%s), reconnect after 10 seconds, retry attempts %v", err.Error(), retryCount)
			<-time.After(10 * time.Second)
			c()
			return
		} else {
			log.Error("AMQP connection failed (%s), close application after 5 retry attempts", err.Error())
			os.Exit(1)

		}
	}

	channelClosed = make(chan *amqp.Error)
	channel.NotifyClose(channelClosed)
	log.Info("AMQP connection established")
	queueList := []string{config.TaskQueueName}
	for _, queueName := range queueList {
		_, err = channel.QueueDeclare(queueName, true, false, false, false, nil)
		if err != nil {
			log.Error("AMQP queue declare failed %s", err.Error())
		}
		delivery, err := consume(queueName)
		if err != nil {
			log.Error("AMQP consume failed %s", err.Error())
		}
		go handleMessage(delivery)
	}
}

func consume(queueName string) (<-chan amqp.Delivery, error) {
	return channel.Consume(queueName, "", false, false, false, false, nil)
}
