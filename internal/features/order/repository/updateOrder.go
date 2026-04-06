package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func UpdateOrder(ctx context.Context, conn *pgx.Conn, orderID int) error {
	sqlQuery := `
    UPDATE orders
	SET status='done'
    WHERE id=$1
`
	_, err := conn.Exec(ctx, sqlQuery, orderID)
	if err != nil {
		return fmt.Errorf("fail update data:%w", err)
	}
	return nil
}

func UpdateOrderStatus(ctx context.Context, conn *pgx.Conn, orderID int, status string) error {
	sqlQuery := `
	UPDATE orders
	SET status=$1
	WHERE id=$2
`
	_, err := conn.Exec(ctx, sqlQuery, status, orderID)
	if err != nil {
		return err
	}
	return nil
}
