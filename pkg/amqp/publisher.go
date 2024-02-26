package amqp

import (
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"judger/pkg/log"
	"time"
)

func Publish(routingKey string, message any) error {
	msg, err := json.Marshal(message)

	log.Debug("Publishing message: %v", string(msg))

	if err == nil {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		err = channel.PublishWithContext(ctx, "", routingKey, false, false, amqp.Publishing{
			ContentType: "application/json",
			Body:        msg,
		})
	}
	return err
}
