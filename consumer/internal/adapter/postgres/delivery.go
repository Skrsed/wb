package pgRepository

import (
	"consumer/internal/core/domain"
	"context"
	"log/slog"
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
		slog.Error("GetDeliveryByUid scanning error", "error", err)
		return nil, err
	}

	return &res, nil
}

func (pr *PostgresRepository) PopulateMapWithDelivery(
	ctx context.Context,
	orders *map[string]*domain.Order,
	uids string,
) error {
	rows, err := pr.db.Query(ctx, "SELECT * from delivery")
	if err != nil {
		slog.Error("PopulateMapWithDelivery error fetching delivery", "error", err)

		return err
	}
	defer rows.Close()

	for rows.Next() {
		var delivery domain.Delivery
		rows.Scan(
			&delivery.ID,
			&delivery.Name,
			&delivery.Phone,
			&delivery.Zip,
			&delivery.City,
			&delivery.Address,
			&delivery.Region,
			&delivery.Email,
			&delivery.OrderUid,
		)

		(*orders)[delivery.OrderUid].Delivery = delivery
	}

	return nil
}
