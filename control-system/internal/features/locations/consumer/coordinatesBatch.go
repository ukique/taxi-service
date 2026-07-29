package consumer

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	models "github.com/ukique/taxi-service/internal/models"
)

func (c *Consumer) CoordinatesBatchConsumer(delivery amqp.Delivery) {

	var event models.OrderCoordinateEvent
	if err := json.Unmarshal(delivery.Body, &event); err != nil {
		log.Print("failed to Unmarshal deliver.Body:", err)
		err := delivery.Nack(false, false)
		if err != nil {
			log.Println("failed to Nack", err)
		}
		return
	}

	// Sending to Coordinates Batch
	c.coordinatesChan <- event
}
