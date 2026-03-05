package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func CreateConnection(ctx context.Context, databaseURL string) *pgx.Conn {
	conn, err := pgx.Connect(ctx, databaseURL)
	if err != nil {
		log.Fatal("fail connect to database:", err)
	}
	return conn
}
