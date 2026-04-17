package ws

import (
	"context"

	"github.com/ukique/taxi-service/internal/features/order/repository"
)

func (h *Handler) OrdersHandler(ctx context.Context, safeConn *safeConn, orderPageID int) {
	// listen inserst/create data in orders
	// it is required that after receiving a change notification, WebSocketHandler sends the current data to client.
	// pgxpool.Pool can't listen
	pgConn, err := h.pool.Acquire(ctx)
	if err != nil {
		return
	}
	defer pgConn.Release()

	if _, err := pgConn.Exec(ctx, "LISTEN orders"); err != nil {
		return
	}

	for {
		if _, err := pgConn.Conn().WaitForNotification(ctx); err != nil {
			return
		}
		orders, err := repository.GetOrdersData(ctx, h.pool, orderPageID)
		if err != nil {
			return
		}
		if err := safeConn.WriteJSON(orders); err != nil {
			return
		}
	}
}
