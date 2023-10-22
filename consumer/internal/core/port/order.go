package port

import (
	"consumer/internal/core/domain"
	"context"
)

type OrderService interface {
	// Order
	GetOrderById(ctx context.Context, orderId int) (*domain.Order, error)
	GetOrdersList(ctx context.Context, limit int) ([]*domain.Order, error)
}
