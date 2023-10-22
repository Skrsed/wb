package service

import (
	cacheRepository "consumer/internal/adapter/cache"
	pgRepository "consumer/internal/adapter/postgres"
	"consumer/internal/core/domain"
	"context"
	"errors"
)

type OrderService struct {
	pg         *pgRepository.PostgresRepository
	orderCache *cacheRepository.OrderCacheRepository
}

func NewOrderService(
	pg *pgRepository.PostgresRepository,
	oc *cacheRepository.OrderCacheRepository,
) (*OrderService, error) {
	return &OrderService{
		pg:         pg,
		orderCache: oc,
	}, nil
}

func (svc *OrderService) GetOrderByUid(ctx context.Context, uid string) (*domain.Order, error) {
	cachedOrder := svc.orderCache.GetOrderByUid(uid)

	if cachedOrder != nil {
		return cachedOrder, nil
	}

	order, errOrder := svc.pg.GetOrderByUid(ctx, uid)
	delivery, errDelivery := svc.pg.GetDeliveryById(ctx, order.DeliveryId)
	payment, errPayment := svc.pg.GetPaymentById(ctx, order.PaymentId)
	items, errItems := svc.pg.GetItemsByOrderUid(ctx, uid)

	if err := errors.Join(errOrder, errDelivery, errPayment, errItems); err != nil {
		return nil, err
	}

	resultOrder := *order
	resultOrder.Delivery = *delivery
	resultOrder.Payment = *payment
	resultOrder.Items = *items

	return &resultOrder, nil
}
