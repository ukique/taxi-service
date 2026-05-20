package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type ConsumerConfig struct {
	QueueName   string
	ConsumerTag string
	AutoAck     bool
	Exclusive   bool
	NoLocal     bool
	NoWait      bool
	Args        amqp.Table
}

func (b *Broker) Consumer(config ConsumerConfig, delivery func(delivery amqp.Delivery)) error {
	messages, err := b.ch.Consume(
		config.QueueName,   // Name
		config.ConsumerTag, // consumer tag
		config.AutoAck,     // auto-a
		config.Exclusive,   // exclusive
		config.NoLocal,     // no-local
		config.NoWait,      // no-wait
		config.Args,        // args
	)
	if err != nil {
		log.Println("failed to consumer order.created: ", err)
		return err
	}
	for message := range messages {
		go delivery(message)
	}
	return nil
}
