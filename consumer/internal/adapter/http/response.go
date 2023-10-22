package httpHandler

import (
	"consumer/internal/core/domain"
)

// ErrorResponse represents an error response body format
type ErrorResponse struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"Error message"`
}

type OrderResponse struct {
	domain.Order
}

// type OrderListResponse struct {
// 	[]domain.Order)
// }
