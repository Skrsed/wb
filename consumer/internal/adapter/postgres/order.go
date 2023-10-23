package pgRepository

import (
	"consumer/internal/core/domain"
	"context"
	"errors"
	"fmt"
	"log/slog"

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
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

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
		slog.Error("Error scanning order result", "error", err)
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

func (pr *PostgresRepository) GetAllOrders(ctx context.Context) (*map[string]*domain.Order, error) {
	rows, err := pr.db.Query(ctx, "SELECT * from orders")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := map[string]*domain.Order{}
	uids := ""

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var order domain.Order
		err := rows.Scan(
			&order.Uid,
			&order.TrackNumber,
			&order.Entry,
			&order.Locale,
			&order.InternalSignature,
			&order.CustomerId,
			&order.DeliveryService,
			&order.Shardkey,
			&order.SmId,
			&order.DateCreated,
			&order.OofShard,
		)

		if err != nil {
			fmt.Println(err)
		}
		order.Items = make([]*domain.Item, 0, 100)
		orders[order.Uid] = &order
		uids += fmt.Sprintf(",'%s'", order.Uid)
	}
	if len(orders) == 0 {
		return nil, nil
	}

	uids = uids[1:]
	pr.PopulateMapWithDelivery(ctx, &orders, uids)
	pr.PopulateMapWithPayments(ctx, &orders, uids)

	pr.PopulateMapWithItems(ctx, &orders, uids)

	return &orders, nil
}
