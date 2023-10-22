package cacheRepository

import (
	"consumer/internal/core/domain"
	"sync"
)

type OrderCacheRepository struct {
	Mutex sync.RWMutex
	Store map[string]domain.Order
}

func NewOrderCacheRepository() *OrderCacheRepository {
	return &OrderCacheRepository{
		Store: make(map[string]domain.Order),
	}
}

func (rep *OrderCacheRepository) SaveOrder(order domain.Order) error {
	rep.Mutex.Lock()
	defer rep.Mutex.Unlock()

	rep.Store[order.Uid] = order

	return nil
}

func (rep *OrderCacheRepository) GetOrderByUid(uid string) *domain.Order {
	rep.Mutex.RLock()
	defer rep.Mutex.RUnlock()

	order, isFound := rep.Store[uid]

	if !isFound {
		return nil
	}

	return &order
}
