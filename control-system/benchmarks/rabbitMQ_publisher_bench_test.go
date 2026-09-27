package benchmarks

import (
	"context"
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/rabbitmq/amqp091-go"
	"github.com/ukique/taxi-service/internal/core/rabbitmq"
	"github.com/ukique/taxi-service/internal/models"
)

var TestOrderCoordinates = &models.OrderCoordinateEvent{
	EventID: 1,
	Order: models.Order{
		ID:       1,
		DriverID: 1,
	},
	Coordinates: models.Coordinates{
		Lat:       47.94467768204419,
		Lon:       132.8670064697572,
		CreatedAt: time.Now(),
	},
}

func BenchmarkProduce_OrderCreated(b *testing.B) {
	err := godotenv.Load("../.env")
	if err != nil {
		b.Fatalf("No .env file, using environment variables")
	}
	rabbitURL := os.Getenv("RABBITMQ_URL_TEST")
	conn, err := amqp091.Dial(rabbitURL)
	if err != nil {
		b.Fatalf("dial error: %v", err)
	}
	b.Cleanup(func() {
		conn.Close()
	})

	ch, err := conn.Channel()
	if err != nil {
		b.Fatalf("channel error: %v", err)
	}
	b.Cleanup(func() {
		ch.Close()
	})

	_, err = ch.QueueDeclare("order.created.test", true, false, false, false, nil)
	if err != nil {
		b.Fatalf("queue declare failed: %v", err)
	}

	b.Cleanup(func() {
		_, err := ch.QueuePurge("order.created.test", false)
		if err != nil {
			b.Logf("failed to purge message from order.created.test: %v", err)
		}
	})

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		messageBody, err := json.Marshal(TestOrderCoordinates)
		if err != nil {
			b.Fatalf("failed to marshal TestOrderCoordinates body %v", err)
		}
		message := amqp091.Publishing{
			Body: messageBody,
		}
		config := rabbitmq.PublisherConfig{
			Exchange:  "",
			Key:       "order.created.test",
			Mandatory: false,
			Immediate: false, // (always false)
			Message:   message,
		}
		err = ch.PublishWithContext(
			context.Background(),
			config.Exchange,
			config.Key,
			config.Mandatory,
			config.Immediate,
			config.Message,
		)
		if err != nil {
			b.Fatalf("publish failed: %v", err)
		}

	}
}

func BenchmarkProduce_OrderCoordinates(b *testing.B) {
	err := godotenv.Load("../.env")
	if err != nil {
		b.Fatalf("No .env file, using environment variables")
	}
	rabbitURL := os.Getenv("RABBITMQ_URL_TEST")
	conn, err := amqp091.Dial(rabbitURL)
	if err != nil {
		b.Fatalf("dial error: %v", err)
	}
	b.Cleanup(func() {
		conn.Close()
	})

	ch, err := conn.Channel()
	if err != nil {
		b.Fatalf("channel error: %v", err)
	}
	b.Cleanup(func() {
		ch.Close()
	})

	_, err = ch.QueueDeclare("order.coordinates.test", true, false, false, false, nil)
	if err != nil {
		b.Fatalf("queue declare failed: %v", err)
	}

	b.Cleanup(func() {
		_, err := ch.QueuePurge("order.coordinates.test", false)
		if err != nil {
			b.Logf("failed to purge message from order.coordinates.test: %v", err)
		}
	})

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		messageBody, err := json.Marshal(TestOrderCoordinates)
		if err != nil {
			b.Fatalf("failed to marshal TestOrderCoordinates body %v", err)
		}
		message := amqp091.Publishing{
			Body: messageBody,
		}

		config := rabbitmq.PublisherConfig{
			Exchange:  "",
			Key:       "order.coordinates.test",
			Mandatory: false,
			Immediate: false, // (always false)
			Message:   message,
		}
		err = ch.PublishWithContext(
			context.Background(),
			config.Exchange,
			config.Key,
			config.Mandatory,
			config.Immediate,
			config.Message,
		)
		if err != nil {
			b.Fatalf("publish failed: %v", err)
		}
	}
}

func BenchmarkProduce_CoordinatesBatch(b *testing.B) {
	err := godotenv.Load("../.env")
	if err != nil {
		b.Fatalf("No .env file, using environment variables")
	}
	rabbitURL := os.Getenv("RABBITMQ_URL_TEST")
	conn, err := amqp091.Dial(rabbitURL)
	if err != nil {
		b.Fatalf("dial error: %v", err)
	}
	b.Cleanup(func() {
		conn.Close()
	})

	ch, err := conn.Channel()
	if err != nil {
		b.Fatalf("channel error: %v", err)
	}
	b.Cleanup(func() {
		ch.Close()
	})

	_, err = ch.QueueDeclare("coordinates.batch.test", true, false, false, false, nil)
	if err != nil {
		b.Fatalf("queue declare failed: %v", err)
	}

	b.Cleanup(func() {
		_, err := ch.QueuePurge("coordinates.batch.test", false)
		if err != nil {
			b.Logf("failed to purge message from coordinates.batch.test: %v", err)
		}
	})

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		messageBody, err := json.Marshal(TestOrderCoordinates)
		if err != nil {
			b.Fatalf("failed to marshal TestOrderCoordinates body %v", err)
		}
		message := amqp091.Publishing{
			Body: messageBody,
		}

		config := rabbitmq.PublisherConfig{
			Exchange:  "",
			Key:       "coordinates.batch.test",
			Mandatory: false,
			Immediate: false, // (always false)
			Message:   message,
		}
		err = ch.PublishWithContext(
			context.Background(),
			config.Exchange,
			config.Key,
			config.Mandatory,
			config.Immediate,
			config.Message,
		)
		if err != nil {
			b.Fatalf("publish failed: %v", err)
		}
	}
}
