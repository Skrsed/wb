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
	// Get a transact repository for making transaction requests.
	prtx, err := pr.StartTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	// Defer a rollback in case anything fails.
	defer prtx.Rollback(ctx)

	order := &domain.Order{}

	items := &[]*domain.Item{}
	errs := []error{}

	order, err = prtx.CreateOrder(ctx, Order)
	if err != nil {
		return nil, err
	}

	order.Payment = Order.Payment
	order.Payment.OrderUid = Order.Uid
	payment, err := prtx.CreatePayment(ctx, &order.Payment)
	errs = append(errs, err)

	order.Delivery = Order.Delivery
	order.Delivery.OrderUid = Order.Uid
	delivery, err := prtx.CreateDelivery(ctx, &order.Delivery)
	errs = append(errs, err)

	items, err = prtx.CreateItems(ctx, Order)
	errs = append(errs, err)

	if err = errors.Join(errs...); err != nil {
		return nil, err
	}

	//Commit the transaction.
	if err = prtx.Commit(ctx); err != nil {
		return nil, err
	}

	order.Payment = *payment
	order.Delivery = *delivery
	order.Items = *items

	// Return the order ID.
	return order, nil
}

// CreateOrder creates a new Order record in the database
// TODO: Try to answer why we're returning pointer insted of actual value? Where that value is? Leaking or not?
func (pr *PostgresRepository) CreateOrder(
	ctx context.Context,
	Order *domain.Order,
) (*domain.Order, error) {
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

func (pr *PostgresRepository) GetOrderByUid(
	ctx context.Context,
	uid string,
) (*domain.Order, error) {
	sql := `SELECT
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
	FROM orders WHERE order_uid = $1`

	fmt.Println("uid - ", uid)

	res := domain.Order{}
	err := pr.db.QueryRow(ctx, sql, uid).Scan(
		&res.Uid,
		&res.TrackNumber,
		&res.Entry,
		&res.Locale,
		&res.InternalSignature,
		&res.CustomerId,
		&res.DeliveryService,
		&res.Shardkey,
		&res.SmId,
		&res.DateCreated,
		&res.OofShard,
	)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &res, nil
}

// func (pr *PostgresRepository) GetAllOrders(ctx context.Context) (*[]*domain.Order, error) {
// 	rows, err := pr.db.Query("SELECT * FROM orders")

// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var items []*domain.Order

// 	// Loop through rows, using Scan to assign column data to struct fields.
// 	for rows.Next() {
// 		var item domain.Order
// 	}
// }
