package order

import (
	"context"
	"encoding/json"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/ukique/taxi-service/config"
	"github.com/ukique/taxi-service/internal/core/rabbitmq"
	"github.com/ukique/taxi-service/internal/features/order/services"
	"github.com/ukique/taxi-service/internal/models"
)

type Consumer struct {
	simulationData config.Config
	broker         *rabbitmq.Broker
	writer         OrdersWriter
}

func NewOrderConsumer(simulationData config.Config, broker *rabbitmq.Broker, writer OrdersWriter) *Consumer {
	return &Consumer{simulationData: simulationData, broker: broker, writer: writer}
}

type OrdersWriter interface {
	UpdateOrderStatus(ctx context.Context, orderID int, status string) error
}

func (c *Consumer) OrderCreatedConsumer(delivery amqp.Delivery) {

	orderCoordinatesPublisherConfig := rabbitmq.PublisherConfig{
		Exchange:  "",
		Key:       "order.coordinates",
		Mandatory: false,
		Immediate: false, // (always false)
	}

	var order models.Order
	if err := json.Unmarshal(delivery.Body, &order); err != nil {
		log.Println("failed to unmarshal delivery.Body:", err)
		err := delivery.Nack(false, false)
		if err != nil {
			log.Println("failed to Nack:", err)
			return
		}
		return
	}
	driverID := order.DriverID
	err := c.writer.UpdateOrderStatus(context.Background(), order.ID, "in_progress")
	if err != nil {
		log.Println("database: fail to change order status:", err)
	}

	var coordinates models.Coordinates
	for i := 1; i <= c.simulationData.Simulator.LocationUpdates; i++ {

		coordinates.Lat, coordinates.Lon, _ = services.GenerateCoordinates()

		event := models.OrderCoordinateEvent{
			EventID: i,
			Coordinates: models.Coordinates{
				Lat: coordinates.Lat,
				Lon: coordinates.Lon,
			},
			Order: models.Order{
				DriverID: driverID,
				ID:       order.ID,
				Status:   order.Status,
			},
		}
		if c.simulationData.Simulator.LocationUpdates == i {
			event.Order.Status = "done"
		}

		orderBody, err := json.Marshal(event)
		if err != nil {
			log.Println("failed to marshal OrderCoordinatesEvent: ", err)
			return
		}

		message := amqp.Publishing{
			Body: orderBody,
		}
		orderCoordinatesPublisherConfig.Message = message
		if err := c.broker.PublisherWithContext(context.Background(), orderCoordinatesPublisherConfig); err != nil {
			log.Println("fail to publish to order.coordinates :", err)
			err := delivery.Nack(false, true)
			if err != nil {
				log.Println("failed to Nack:", err)
				return
			}
			return
		}
		time.Sleep(c.simulationData.Simulator.TimeOut)
	}
	if err := delivery.Ack(false); err != nil {
		log.Println("failed to send Ack message:", err)
	}
}
