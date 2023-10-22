package pgRepository

import (
	"consumer/internal/core/domain"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (pr *PostgresRepository) CreateOrderCascade(
	ctx context.Context,
	Order *domain.Order,
) (*domain.Order, error) {
	// Create a helper function for preparing failure results.
	fail := func(err error) (*domain.Order, error) {
		return nil, fmt.Errorf("CreateOrder: %v", err)
	}

	// Get a Tx for making transaction requests.
	// TODO: Check avalivable options
	tx, err := pr.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fail(err)
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback(ctx)

	order, e1 := pr.CreateOrder(ctx, Order)
	payment, e2 := pr.CreatePayment(ctx, &Order.Payment)
	if err != nil {
		return fail(err)
	}
	delivery, e3 := pr.CreateDelivery(ctx, &Order.Delivery)
	if err != nil {
		return fail(err)
	}
	items, e4 := pr.CreateItems(ctx, &Order.Items)
	if err != nil {
		return fail(err)
	}

	// Commit the transaction.
	if err := errors.Join(tx.Commit(ctx), e1, e2, e3, e4); err != nil {
		return fail(err)
	}

	order.Payment = *payment
	order.Delivery = *delivery
	order.Items = *items

	// Return the order ID.
	return order, nil
}

// CreateOrder creates a new Order record in the database
// TODO: Try to answer why we're returning pointer insted of actual value? Where that value is? Leaking or not?
func (pr *PostgresRepository) CreateOrder(ctx context.Context, Order *domain.Order) (*domain.Order, error) {
	sql := `INSERT INTO orders (
		order_uid,
		track_number,
		entry,
		locale,
		internal_signature,
		customer_id,
		delivery_service,
		shardkey,
		sm_id,
		date_created,
		oof_shard
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	RETURNING order_uid` // test perfomance

	res := *Order
	err := pr.db.QueryRow(ctx, sql,
		&Order.Uid,
		&Order.TrackNumber,
		&Order.Entry,
		&Order.Locale,
		&Order.InternalSignature,
		&Order.CustomerId,
		&Order.DeliveryService,
		&Order.Shardkey,
		&Order.SmId,
		&Order.DateCreated,
		&Order.OofShard,
	).Scan(&res.Uid)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
