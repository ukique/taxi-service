package rabbitmq

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type QueueConfig struct {
	Name       string
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
	Args       amqp.Table
}

func (b *Broker) DeclareQueue(config QueueConfig) (amqp.Queue, error) {
	queue, err := b.ch.QueueDeclare(
		config.Name,       //name
		config.Durable,    // durable
		config.AutoDelete, // autoDelete
		config.Exclusive,  // exclusive
		config.NoWait,     //noWait
		config.Args,       //args
	)
	if err != nil {
		return amqp.Queue{}, fmt.Errorf("failed to declare queue %s: %w", config.Name, err)
	}
	return queue, nil
}
