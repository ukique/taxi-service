package repository

import "github.com/jackc/pgx/v5/pgxpool"

type DriversRepository struct {
	pool *pgxpool.Pool
}

func NewDriversRepository(pool *pgxpool.Pool) *DriversRepository {
	return &DriversRepository{pool: pool}
}
