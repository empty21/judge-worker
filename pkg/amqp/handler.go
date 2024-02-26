package amqp

import (
	"encoding/json"
	"errors"
	amqp "github.com/rabbitmq/amqp091-go"
	"judger/pkg/config"
	"judger/pkg/judger"
	"judger/pkg/log"
	"judger/pkg/model"
)

func handleMessage(messages <-chan amqp.Delivery) {
	for message := range messages {
		log.Info("Received message: %s", message.Body)
		switch message.RoutingKey {
		case config.TaskQueueName:
			_ = handleJudgeMessage(message.Body)
			break
		default:
			_ = errors.New("no handler registered")
		}
		_ = message.Ack(false)
	}
}

func handleJudgeMessage(message []byte) error {
	var task = &model.JudgeTask{}
	err := json.Unmarshal(message, &task)
	if err != nil {
		return err
	}
	err = Publish(config.ResultQueueName, model.NewJudgeTaskStatus(task.Identifier, config.TaskStatusIP, ""))
	if err != nil {
		return err
	}
	result := judger.Judger.Judge(*task)
	err = Publish(config.ResultQueueName, result)
	return err
}
