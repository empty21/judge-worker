package amqp

import (
	"encoding/json"
	"errors"
	"github.com/streadway/amqp"
	"judger/pkg/domain"
	"judger/pkg/judger"
)

func handleMessage(message amqp.Delivery) error {
	var msg Message
	err := json.Unmarshal(message.Body, &msg)
	if err != nil {
		return err
	}
	switch msg.Pattern {
	case JudgeTaskNew:
		err = handleJudgeMessage(message.Body)
		break
	default:
		return errors.New("no handler registered")
	}
	return err
}

func handleJudgeMessage(message []byte) error {
	var msg Message
	var data = &domain.JudgeTask{}
	msg.Data = data
	err := json.Unmarshal(message, &msg)
	if err != nil {
		return err
	}
	err = Publish(JudgeTaskUpdate, Message{
		Pattern: JudgeTaskUpdate,
		Data:    domain.NewJudgeTaskStatus(data.Uid, domain.TaskStatusIP),
	})
	if err != nil {
		return err
	}
	result := judger.Judger.Judge(*data)
	err = Publish(JudgeTaskUpdate, Message{
		Pattern: JudgeTaskUpdate,
		Data:    result,
	})
	if err != nil {
		return err
	}
	return nil
}
