package port

import (
	"consumer/internal/core/domain"
	"context"
)

// PaymentRepository is an interface for interacting with payment-related data
type PaymentRepository interface {
	// CreatePayment inserts a new payment into the database
	CreatePayment(ctx context.Context, payment *domain.Payment) (*domain.Payment, error)
	// UpdatePayment updates a payment
	UpdatePayment(ctx context.Context, payment *domain.Payment) (*domain.Payment, error)
}

// PaymentService is an interface for interacting with payment-related business logic
// type PaymentService interface {
// 	// CreatePayment creates a new payment
// 	CreatePayment(ctx context.Context, payment *domain.Payment) (*domain.Payment, error)
// 	// UpdatePayment updates a payment
// 	UpdatePayment(ctx context.Context, payment *domain.Payment) (*domain.Payment, error)
// }
