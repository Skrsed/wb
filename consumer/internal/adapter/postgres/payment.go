package pgRepository

import (
	"consumer/internal/core/domain"
	"context"
)

// CreatePayment creates a new payment record in the database
// TODO: Try to answer why we're returning pointer insted of actual value? Where that value is? Leaking or not?
func (pr *PostgresRepository) CreatePayment(
	ctx context.Context,
	Payment *domain.Payment,
) (*domain.Payment, error) {
	sql := `INSERT INTO payments (
		transaction,
		request_id,
		currency,
		provider,
		amount,
		payment_dt,
		bank,
		delivery_cost,
		goods_total,
		custom_fee,
		order_uid
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	RETURNING *` // test perfomance

	res := *Payment
	err := pr.db.QueryRow(ctx, sql,
		&Payment.Transaction,
		&Payment.RequestId,
		&Payment.Currency,
		&Payment.Provider,
		&Payment.Amount,
		&Payment.PaymentDt,
		&Payment.Bank,
		&Payment.DeliveryCost,
		&Payment.GoodsTotal,
		&Payment.CustomFee,
		&Payment.OrderUid,
	).Scan(
		&res.ID,
		&res.Transaction,
		&res.RequestId,
		&res.Currency,
		&res.Provider,
		&res.Amount,
		&res.PaymentDt,
		&res.Bank,
		&res.DeliveryCost,
		&res.GoodsTotal,
		&res.CustomFee,
		&res.OrderUid,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (pr *PostgresRepository) GetPaymentByUid(
	ctx context.Context,
	uid string,
) (*domain.Payment, error) {
	sql := `SELECT
		id,
		transaction,
		request_id,
		currency,
		provider,
		amount,
		payment_dt,
		bank,
		delivery_cost,
		goods_total,
		custom_fee,
		order_uid
	FROM payments WHERE order_uid = $1`

	res := domain.Payment{}
	err := pr.db.QueryRow(ctx, sql, uid).Scan(
		&res.ID,
		&res.Transaction,
		&res.RequestId,
		&res.Currency,
		&res.Provider,
		&res.Amount,
		&res.PaymentDt,
		&res.Bank,
		&res.DeliveryCost,
		&res.GoodsTotal,
		&res.CustomFee,
		&res.OrderUid,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
