package rabbitmq

import (
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Broker struct {
	conn *amqp.Connection
	url  string
	ch   *amqp.Channel
}

func NewBroker(url string) (*Broker, error) {
	var conn *amqp.Connection
	var err error

	for i := 0; i < 10; i++ {
		conn, err = amqp.Dial(url)
		if err == nil {
			break
		}
		log.Printf("failed to connect RabbitMQ, retry %d/10: %v", i+1, err)
		time.Sleep(time.Second * 3)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to connect RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to create RabbitMQ channel: %w", err)
	}
	return &Broker{
		conn: conn,
		url:  url,
		ch:   ch,
	}, nil
}
func (b *Broker) Close() error {
	if b.ch != nil {
		if err := b.ch.Close(); err != nil {
			return fmt.Errorf("failed to close RabbitMQ channel: %w", err)
		}
	}
	if b.conn != nil {
		if err := b.conn.Close(); err != nil {
			return fmt.Errorf("failed to close RabbitMQ connection: %w", err)
		}
	}
	return nil
}
