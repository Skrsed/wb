package port

import (
	"consumer/internal/core/domain"
	"context"
)

// OrderRepository is an interface for interacting with Order-related data
type OrderRepository interface {
	// CreateOrder inserts a new Order into the database
	CreateOrder(ctx context.Context, Order *domain.Order) (*domain.Order, error)
	// UpdateOrder updates a Order
	UpdateOrder(ctx context.Context, Order *domain.Order) (*domain.Order, error)
}

// OrderService is an interface for interacting with Order-related business logic
// type OrderService interface {
// 	// CreateOrder creates a new Order
// 	CreateOrder(ctx context.Context, Order *domain.Order) (*domain.Order, error)
// 	// UpdateOrder updates a Order
// 	UpdateOrder(ctx context.Context, Order *domain.Order) (*domain.Order, error)
// }
