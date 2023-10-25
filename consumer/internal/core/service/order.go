package service

import (
	cacheRepository "consumer/internal/adapter/cache"
	pgRepository "consumer/internal/adapter/postgres"
	"consumer/internal/core/domain"
	"context"
	"errors"
	"log/slog"

	"github.com/go-playground/validator/v10"
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
		slog.Info("row was finded in cache")
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
	validate := validator.New()
	err := validate.Struct(order)
	if err != nil {
		slog.Error("SaveOrder error validating order", "error", err)
		return err
	}

	newOrder, err := svc.pg.CreateOrderCascade(ctx, order)
	if err != nil {
		slog.Error("SaveOrder error saving db", "error", err)
		return err
	}

	err = svc.orderCache.SaveOrder(newOrder)
	if err != nil {
		slog.Error("SaveOrder error saving cache", "error", err)
		return err
	}

	return nil
}

func (svc *OrderService) LoadCacheFromDb(ctx context.Context) error {
	orders, err := svc.pg.GetAllOrders(ctx)
	if err != nil {
		return err
	}
	if orders == nil {
		return nil
	}

	svc.orderCache.PutAll(orders)

	return nil
}
