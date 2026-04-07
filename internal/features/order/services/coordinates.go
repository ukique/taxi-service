package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"

	"github.com/jackc/pgx/v5"
	amqp "github.com/rabbitmq/amqp091-go"
	orderFeatures "github.com/ukique/taxi-service/internal/features/order/repository"
	"github.com/ukique/taxi-service/internal/models"
)

// GenerateCoordinates generate random float64 coordinates
// for example lat = 47.842658, lon = 34.811989
func GenerateCoordinates() (float64, float64) {
	driverLat := rand.Float64()*180 - 90
	driverLon := rand.Float64()*360 - 180
	return driverLat, driverLon
}

// SendCoordinates sends to the message broker messages of the following structure
// Example structure you can check in docs/examples/message_broker_structure.md
func SendCoordinates(ctx context.Context, conn *pgx.Conn, ch *amqp.Channel, orderCoordinatesQueue amqp.Queue, orders []models.Order) error {
	for _, order := range orders {
		//Getting random coordinates
		var coordinates models.Coordinates
		coordinates.Lat, coordinates.Lon = GenerateCoordinates()

		order.Status = "in_progress"
		orderStatus := "in_progress"
		err := orderFeatures.UpdateOrderStatus(ctx, conn, order.ID, orderStatus)
		if err != nil {
			fmt.Println("fail to change order status:", err)
			return err
		}

		//collect message for message broker
		orderBody := models.OrderCoordinateEvent{
			DriverID: order.DriverID,
			Coordinates: models.Coordinates{
				Lat: coordinates.Lat,
				Lon: coordinates.Lon,
			},
			Order: models.Order{
				ID:     order.ID,
				Status: order.Status,
			},
		}
		orderData, err := json.Marshal(orderBody)
		if err != nil {
			log.Println("fail to marshal orderBody:")
			return err
		}

		err = ch.Publish(
			"", //exchange
			orderCoordinatesQueue.Name,
			false, // mandatory
			false, // immediate
			amqp.Publishing{
				ContentType: "application/json",
				Body:        orderData,
			},
		)
		if err != nil {
			log.Println("fail to publish orderData:", err)
			return err
		}
	}
	return nil
}
