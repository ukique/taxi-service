package connections

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/ukique/taxi-service/internal/core/database"
	"github.com/ukique/taxi-service/internal/core/rabbitmq"
)

type Connections struct {
	Pool      *pgxpool.Pool
	Broker    *rabbitmq.Broker
	SecretKey string
}

func LoadConnections() *Connections {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file", err)
		os.Exit(1)
	}
	// get DATABASE_URL from .env
	dataBaseURL := os.Getenv("DATABASE_URL")

	//get SECRET_KEY for JWT
	secretKey := os.Getenv("SECRET_KEY")

	//get RABBITMQ_URL from .env
	rabbitmqURL := os.Getenv("RABBITMQ_URL")

	// create *Conn for database features
	ctx := context.Background()
	pool, err := database.CreateConnection(ctx, dataBaseURL)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	broker, err := rabbitmq.NewBroker(rabbitmqURL)
	if err != nil {
		log.Println("failed to create NewBroker:", err)
	}

	return &Connections{Pool: pool, Broker: broker, SecretKey: secretKey}
}
