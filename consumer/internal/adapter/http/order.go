package httpHandler

import (
	"consumer/internal/core/port"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	svc port.OrderService
}

// NewOrderHandler creates a new OrderHandler instance
func NewOrderHandler(svc port.OrderService) *OrderHandler {
	return &OrderHandler{
		svc,
	}
}

// getOrderRequest represents a request body for retrieving an order
type getOrderRequest struct {
	UID string `uri:"uid" binding:"required" example:"b563feb7b2b84b6test"`
}

// GetOrderById godoc
// @Summary      Get an order by ID
// @Description  Get order with related entities by order_uid
// @Tags         Order
// @Accept       json
// @Produce      json
// @Param        uid   path      string  true  "Order UID"
// @Success      200  {object}  OrderResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /v1/order/{uid} [get]
func (h *OrderHandler) GetOrderByUId(ctx *gin.Context) {
	var req getOrderRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		fmt.Println(err)
		validationError(ctx, err)
		return
	}

	order, err := h.svc.GetOrderByUid(ctx, req.UID)
	if err != nil {
		fmt.Println(err)
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, OrderResponse{
		order,
	})
}

func (h *OrderHandler) GetListOrders(ctx *gin.Context) {}
