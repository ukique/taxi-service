package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ukique/taxi-service/internal/models"
)

func GetAllDrivers(ctx context.Context, pool *pgxpool.Pool) ([]models.Driver, error) {
	sqlQuery := `
  SELECT id,username, status FROM drivers;
`
	var drivers []models.Driver
	rows, err := pool.Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var d models.Driver
		if err := rows.Scan(&d.ID, &d.Username, &d.Status); err != nil {
			return nil, err
		}
		drivers = append(drivers, d)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return drivers, nil
}
