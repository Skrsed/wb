package pgRepository

import (
	"consumer/internal/core/domain"
	"context"
	"fmt"
)

// CreateDelivery creates a new Delivery record in the database
// TODO: Try to answer why we're returning pointer insted of actual value? Where that value is? Leaking or not?
func (pr *PostgresRepository) CreateDelivery(
	ctx context.Context,
	Delivery *domain.Delivery,
) (*domain.Delivery, error) {
	sql := `INSERT INTO delivery (
		name,
		phone,
		zip,
		city,
		address,
		region,
		email,
		order_uid
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING *` // TODO: WHAT WE SHOLD DO ON DUPLICATE ORDER ID? CASCADE UPSERT? NOTHING?

	res := *Delivery
	err := pr.db.QueryRow(ctx, sql,
		&Delivery.Name,
		&Delivery.Phone,
		&Delivery.Zip,
		&Delivery.City,
		&Delivery.Address,
		&Delivery.Region,
		&Delivery.Email,
		&Delivery.OrderUid,
	).Scan(
		&res.ID,
		&Delivery.Name,
		&Delivery.Phone,
		&Delivery.Zip,
		&Delivery.City,
		&Delivery.Address,
		&Delivery.Region,
		&Delivery.Email,
		&Delivery.OrderUid,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (pr *PostgresRepository) GetDeliveryByUid(
	ctx context.Context,
	uid string,
) (*domain.Delivery, error) {
	// TODO: extend with fields
	sql := `SELECT
		id,
		name,
		phone,
		zip,
		city,
		address,
		region,
		email,
		order_uid
	FROM delivery WHERE order_uid = $1`

	res := domain.Delivery{}
	err := pr.db.QueryRow(ctx, sql, uid).Scan(
		&res.ID,
		&res.Name,
		&res.Phone,
		&res.Zip,
		&res.City,
		&res.Address,
		&res.Region,
		&res.Email,
		&res.OrderUid,
	)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &res, nil
}
