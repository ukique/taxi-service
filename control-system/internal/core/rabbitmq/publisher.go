package rabbitmq

import (
	"context"
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

type PublisherConfig struct {
	Exchange  string
	Key       string
	Mandatory bool
	Immediate bool // (always false)
	Message   amqp091.Publishing
}

func (b *Broker) PublisherWithContext(ctx context.Context, config PublisherConfig) error {
	err := b.ch.PublishWithContext(
		ctx,
		config.Exchange,
		config.Key,
		config.Mandatory,
		config.Immediate,
		config.Message,
	)
	if err != nil {
		return fmt.Errorf("failed to publish %s: %w", config.Key, err)
	}
	return nil
}
