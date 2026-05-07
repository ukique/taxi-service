package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ukique/taxi-service/internal/models"
)

type DriversRepository struct {
	pool *pgxpool.Pool
}

func NewDriversRepository(pool *pgxpool.Pool) *DriversRepository {
	return &DriversRepository{pool: pool}
}
func (d DriversRepository) GetDriversData(ctx context.Context, pageID int) ([]models.Driver, error) {
	sqlQuery := `
  SELECT id,username, status FROM drivers
  ORDER BY id DESC 
  LIMIT $1 OFFSET $2;
`
	recordsLimit := 50
	offest := recordsLimit * (pageID - 1)
	var drivers []models.Driver
	rows, err := d.pool.Query(ctx, sqlQuery, recordsLimit, offest)
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
