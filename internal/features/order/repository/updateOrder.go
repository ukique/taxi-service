package repository

import (
	"context"
	"fmt"
)

func (o *OrderRepository) UpdateOrder(ctx context.Context, orderID int) error {
	sqlQuery := `
    UPDATE orders
	SET status='done'
    WHERE id=$1
`
	_, err := o.pool.Exec(ctx, sqlQuery, orderID)
	if err != nil {
		return fmt.Errorf("fail update data:%w", err)
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
