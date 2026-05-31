package repository

import (
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DriversRepository struct {
	pool *pgxpool.Pool
}

func NewDriversRepository(pool *pgxpool.Pool) *DriversRepository {
	return &DriversRepository{pool: pool}
}

var ErrNoDriverAvailable = errors.New("no available driver")
