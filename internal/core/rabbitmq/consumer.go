package rabbitmq

import (
	"context"
	"encoding/json"
	"log"

	"github.com/jackc/pgx/v5"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/ukique/taxi-service/internal/features/locations/repository"
	"github.com/ukique/taxi-service/internal/models"
)

func LocationDatabaseConsumer(ctx context.Context, conn *pgx.Conn, ch *amqp.Channel, orderCoordinatesQueue amqp.Queue) error {
	orders, err := ch.Consume(
		orderCoordinatesQueue.Name,
		"",    // consumer tag
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		log.Println("fail to ")
		return err
	}
	for order := range orders {
		var event models.OrderCoordinateEvent
		if err := json.Unmarshal(order.Body, &event); err != nil {
			log.Println("failed to unmarshal orderBody: ", err)
			return err
		}
		log.Println("saving location for driver:", event.DriverID)
		go repository.SaveLocation(ctx, conn, event)
	}
	return nil
}
