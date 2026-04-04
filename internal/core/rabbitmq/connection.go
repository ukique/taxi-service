package rabbitmq

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

// CreateRabbitMQConnection create a RabbitMQ Connection
func CreateRabbitMQConnection(rabbitmqURL string) (*amqp.Connection, error) {
	conn, err := amqp.Dial(rabbitmqURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect RabbitMQ: %w", err)
	}
	return conn, nil
}

// CreateChannel create a channel for working with RabbitMQ
func CreateChannel(brokerConn *amqp.Connection) (*amqp.Channel, error) {
	ch, err := brokerConn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to create rabbitMQ channel: %w", err)
	}
	return ch, nil
}

// QueueDeclareOrdersCoordinates create orders_coordinates queue
func QueueDeclareOrdersCoordinates(ch *amqp.Channel) (amqp.Queue, error) {
	queue, err := ch.QueueDeclare(
		"orders_coordinates", // name
		false,                // durable
		false,                // autoDelete
		false,                // exclusive
		false,                //noWait
		nil,                  //args
	)
	if err != nil {
		return amqp.Queue{}, fmt.Errorf("failed to declare a queue: %w", err)
	}
	return queue, nil
}
