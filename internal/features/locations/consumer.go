package locations

import (
	"context"
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/ukique/taxi-service/internal/features/order/transport"
	"github.com/ukique/taxi-service/internal/models"
)

type Consumer struct {
	hub    transport.Broadcaster
	writer LocationWriter
}

func NewLocationConsumer(writer LocationWriter, hub transport.Broadcaster) *Consumer {
	return &Consumer{writer: writer, hub: hub}
}

type LocationWriter interface {
	SaveLocation(ctx context.Context, orderBody models.OrderCoordinateEvent) error
}

func (c *Consumer) OrderLocationConsumer(delivery amqp.Delivery) {

	var eventBody models.OrderCoordinateEvent
	if err := json.Unmarshal(delivery.Body, &eventBody); err != nil {
		log.Println("failed to unmarshal delivery.Body:", err)
		err := delivery.Nack(false, false)
		if err != nil {
			log.Println("failed to Nack:", err)
			return
		}
		return
	}

	//Saving to driver_locations table
	if err := c.writer.SaveLocation(context.Background(), eventBody); err != nil {
		log.Println("failed to SaveLocation: ", err)
		err := delivery.Nack(false, true)
		if err != nil {
			log.Println("failed to Nack:", err)
			return
		}
		return
	}

	// Sending to BroadCast where subscribe_orderDetails
	messageBody := models.OutgoingMessage[models.OrderCoordinateEvent]{
		Type: "coordinates",
		Page: eventBody.Order.ID,
		Data: eventBody,
	}

	message, err := json.Marshal(messageBody)
	if err != nil {
		log.Println("failed marshal coordinates OutGoingMessage:", err)
		err := delivery.Nack(false, true)
		if err != nil {
			log.Println("failed to Nack:", err)
			return
		}
		return
	}
	c.hub.SendToBroadcast(message)

	if err := delivery.Ack(false); err != nil {
		log.Println("failed to send Ack message:", err)
	}
}
