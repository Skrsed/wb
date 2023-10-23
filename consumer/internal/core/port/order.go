package port

import (
	"consumer/internal/core/domain"
	"context"
)

type OrderService interface {
	// Order
	GetOrderByUid(ctx context.Context, uid string) (*domain.Order, error)
	SaveOrder(ctx context.Context, order *domain.Order) error
	//GetOrdersList(ctx context.Context, limit int) ([]*domain.Order, error)
}
