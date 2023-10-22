package httpHandler

import (
	"consumer/internal/core/port"
	svc "consumer/internal/core/port"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	svc.OrderService
}

// NewOrderHandler creates a new OrderHandler instance
func NewOrderHandler(svc port.OrderService) *OrderHandler {
	return &OrderHandler{
		svc,
	}
}

// GetOrderById godoc
// @Summary      Get an order by ID
// @Description  Get order with related entities by order_uid
// @Tags         Order
// @Accept       json
// @Produce      json
// @Param        order_uid   path      string  true  "Order UID"
// @Success      200  {object}  OrderResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /order/{uid} [get]
func (h *OrderHandler) GetOrderById(ctx *gin.Context) {

}

func (h *OrderHandler) GetListOrders(ctx *gin.Context) {
}
