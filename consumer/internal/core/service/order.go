package service

import (
	cacheRepository "consumer/internal/adapter/cache"
	pgRepository "consumer/internal/adapter/postgres"
	"consumer/internal/core/domain"
	"context"
	"errors"
	"fmt"
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
		fmt.Println("goted from cache", cachedOrder)
		return cachedOrder, nil
	}

	order, err := svc.pg.GetOrderByUid(ctx, uid)
	if err != nil {
		return nil, err
	}

	errs := []error{}
	delivery, err := svc.pg.GetDeliveryByUid(ctx, order.Uid)
	errs = append(errs, err)
	payment, err := svc.pg.GetPaymentByUid(ctx, order.Uid)
	errs = append(errs, err)
	items, err := svc.pg.GetItemsByOrderUid(ctx, uid, order.TrackNumber)
	errs = append(errs, err)

	if err := errors.Join(errs...); err != nil {
		return nil, err
	}

	resultOrder := *order
	resultOrder.Delivery = *delivery
	resultOrder.Payment = *payment
	resultOrder.Items = *items

	return &resultOrder, nil
}

func (svc *OrderService) SaveOrder(ctx context.Context, order *domain.Order) error {
	order, err := svc.pg.CreateOrderCascade(ctx, order)
	if err != nil {
		return err
	}

	err = svc.orderCache.SaveOrder(*order)
	if err != nil {
		fmt.Println("Error while saving in cache", err)
		return err
	}

	return nil
}

// func (svc *OrderService) LoadCacheFromDb(ctx context.Context) error {
// 	orders, err := svc.pg.GetAllOrders(ctx)
// }
