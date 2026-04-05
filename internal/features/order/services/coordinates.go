package services

import (
	"encoding/json"
	"fmt"
	"math/rand"

	amqp "github.com/rabbitmq/amqp091-go"
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
func SendCoordinates(ch *amqp.Channel, orderCoordinatesQueue amqp.Queue, order models.Order) error {

	//Getting random coordinates
	var coordinates models.Coordinates
	coordinates.Lat, coordinates.Lon = GenerateCoordinates()

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
		return fmt.Errorf("fail to marshal orderBody: %w", err)
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
		return fmt.Errorf("fail to publish orderData: %w", err)
	}
	return nil
}
