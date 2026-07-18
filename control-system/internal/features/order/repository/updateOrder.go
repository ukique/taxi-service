package repository

import (
	"context"
	"time"
)

func (o *OrderRepository) UpdateOrder(ctx context.Context, orderID int) error {
	sqlQuery := `
    UPDATE orders
	SET status='done', finished_at=$1
    WHERE id=$2
`
	_, err := o.pool.Exec(ctx, sqlQuery, time.Now(), orderID)
	if err != nil {
		return err
	}
	return nil
}

func (o *OrderRepository) UpdateOrderStatus(ctx context.Context, orderID int, status string) error {
	sqlQuery := `
	UPDATE orders
	SET status=$1
	WHERE id=$2
`
	_, err := o.pool.Exec(ctx, sqlQuery, status, orderID)
	if err != nil {
		return err
	}
	return nil
}
