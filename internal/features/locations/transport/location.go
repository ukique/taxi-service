package transport

import (
	"context"

	"github.com/ukique/taxi-service/internal/models"
)

type LocationHistoryGetter interface {
	GetOrderLocationHistory(ctx context.Context, orderID int) ([]models.OrderCoordinateEvent, error)
}
type Handler struct {
	getter    LocationHistoryGetter
	secretKey string
}

func NewLocationHandler(getter LocationHistoryGetter, secretKey string) *Handler {
	return &Handler{getter: getter, secretKey: secretKey}
}
