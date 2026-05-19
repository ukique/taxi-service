package locations

import (
	"context"
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/ukique/taxi-service/internal/models"
)

type Consumer struct {
	writer LocationWriter
}

func NewLocationConsumer(writer LocationWriter) *Consumer {
	return &Consumer{writer: writer}
}

type LocationWriter interface {
	SaveLocation(ctx context.Context, orderBody models.OrderCoordinateEvent) error
}

func (c *Consumer) OrderLocationConsumer(delivery amqp.Delivery) {

	var event models.OrderCoordinateEvent
	if err := json.Unmarshal(delivery.Body, &event); err != nil {
		log.Println("failed to unmarshal delivery.Body:", err)
		err := delivery.Nack(false, false)
		if err != nil {
			log.Println("failed to Nack:", err)
			return
		}
		return
	}

	//Saving to driver_locations table
	if err := c.writer.SaveLocation(context.Background(), event); err != nil {
		log.Println("failed to SaveLocation: ", err)
		err := delivery.Nack(false, true)
		if err != nil {
			log.Println("failed to Nack:", err)
			return
		}
		return
	}
	// Sending to BroadCast where subscribe_orderDetails with event.Order.ID
	if err := delivery.Ack(false); err != nil {
		log.Println("failed to send Ack message:", err)
	}
}
