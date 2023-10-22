package pgRepository

import (
	"consumer/internal/core/domain"
	"context"
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
	RETURNING id` // test perfomance

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
	).Scan(&res.ID)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
