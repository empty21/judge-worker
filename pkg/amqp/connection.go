package amqp

import amqp "github.com/rabbitmq/amqp091-go"

func connect(uri string) (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial(uri)
	if err != nil {
		return nil, nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, err
	}
	_ = ch.Qos(1, 0, false)
	return conn, ch, nil
}
