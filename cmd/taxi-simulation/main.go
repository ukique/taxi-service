package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/goccy/go-yaml"
	"github.com/joho/godotenv"
	"github.com/rabbitmq/amqp091-go"
	"github.com/ukique/taxi-service/config"
	"github.com/ukique/taxi-service/internal/core/database"
	"github.com/ukique/taxi-service/internal/core/rabbitmq"
	orderFeatures "github.com/ukique/taxi-service/internal/features/order/repository"
	"github.com/ukique/taxi-service/internal/features/order/services"
	"github.com/ukique/taxi-service/internal/models"
)

func main() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file", err)
		os.Exit(1)
	}

	// pool must be replaced (after refactor UpdateOrderStatus)
	dataBaseURL := os.Getenv("DATABASE_URL")
	pool, err := database.CreateConnection(context.Background(), dataBaseURL)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	//get RABBITMQ_URL from .env
	rabbitmqURL := os.Getenv("RABBITMQ_URL")

	//Get LoadTesting Simulation Data
	var simulationData config.Config

	file, err := os.ReadFile("config/simulation.yaml")
	if err != nil {
		log.Println("failed to load simulation.yaml:", err)
		os.Exit(1)
	}
	if err := yaml.Unmarshal(file, &simulationData); err != nil {
		log.Println("failed to unmarshal simulationData:", err)
		os.Exit(1)
	}
	broker, err := rabbitmq.NewBroker(rabbitmqURL)
	if err != nil {
		log.Println("failed to create NewBroker:", err)
	}
	defer func() {
		if err := broker.Close(); err != nil {
			log.Printf("failed to close broker: %v", err)
		}
	}()

	orderCoordinatesQueueConfig := rabbitmq.QueueConfig{
		Name:       "order.coordinates",
		Durable:    true,
		AutoDelete: false,
		Exclusive:  false,
		NoWait:     false,
		Args:       nil,
	}
	_, err = broker.DeclareQueue(orderCoordinatesQueueConfig)
	if err != nil {
		log.Println("fail to Declare Queue order.coordinates :", err)
		os.Exit(1)
	}
	orderCreatedConsumerConfig := rabbitmq.ConsumerConfig{
		QueueName:   "order.created",
		ConsumerTag: "",
		AutoAck:     false,
		Exclusive:   false,
		NoLocal:     false,
		NoWait:      false,
		Args:        nil,
	}
	for {
		err = broker.Consumer(orderCreatedConsumerConfig, func(delivery amqp091.Delivery) {

			orderCoordinatesPublisherConfig := rabbitmq.PublisherConfig{
				Exchange:  "",
				Key:       "order.coordinates",
				Mandatory: false,
				Immediate: false, // (always false)
			}

			var order models.Order
			if err := json.Unmarshal(delivery.Body, &order); err != nil {
				log.Println("failed to unmarshal delivery.Body:", err)
				delivery.Nack(false, false)
				return
			}
			driverID := order.DriverID
			order.Status = "in_progress"
			err := orderFeatures.UpdateOrderStatus(context.Background(), pool, order.ID, "in_progress")
			if err != nil {
				log.Println("database: fail to change order status:", err)
			}

			var coordinates models.Coordinates
			for i := 0; i < simulationData.Simulator.LocationUpdates; i++ {

				coordinates.Lat, coordinates.Lon, _ = services.GenerateCoordinates()

				event := models.OrderCoordinateEvent{
					DriverID: driverID,
					Coordinates: models.Coordinates{
						Lat: coordinates.Lat,
						Lon: coordinates.Lon,
					},
					Order: models.Order{
						ID:     order.ID,
						Status: order.Status,
					},
				}
				log.Println("coordinates", coordinates.Lon, "", coordinates.Lat)
				orderBody, err := json.Marshal(event)
				if err != nil {
					log.Println("failed to marshal OrderCoordinatesEvent: ", err)
					return
				}

				message := amqp091.Publishing{
					Body: orderBody,
				}
				orderCoordinatesPublisherConfig.Message = message
				if err := broker.PublisherWithContext(context.Background(), orderCoordinatesPublisherConfig); err != nil {
					log.Println("fail to publish to order.coordinates :", err)
					delivery.Nack(false, true)
					return
				}
				time.Sleep(simulationData.Simulator.TimeOut)
			}
			if err := delivery.Ack(false); err != nil {
				log.Println("failed to send Ack message:", err)
			}
		})
		if err != nil {
			log.Println("failed to consume order.created:", err)
		}

	}
}
