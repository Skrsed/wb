package repository

import (
	"consumer/internal/core/domain"
	"context"
)

/**
 * PaymentRepository implements port.PaymentRepository interface
 * and provides an access to the postgres database
 */
type PaymentRepository struct {
	db *DB
}

// NewPaymentRepository creates a new payment repository instance
func NewPaymentRepository(db *DB) *PaymentRepository {
	return &PaymentRepository{
		db,
	}
}

// CreatePayment creates a new payment record in the database
// TODO: Try to answer why we're returning pointer insted of actual value? Where that value is? Leaking or not?
func (pr *PaymentRepository) CreatePayment(ctx context.Context, payment *domain.Payment) (*domain.Payment, error) {
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
		custom_fee
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	RETURNING id` // test perfomance

	res := *payment
	err := pr.db.QueryRow(ctx, sql,
		&payment.Transaction,
		&payment.RequestId,
		&payment.Currency,
		&payment.Provider,
		&payment.Amount,
		&payment.PaymentDt,
		&payment.Bank,
		&payment.DeliveryCost,
		&payment.GoodsTotal,
		&payment.CustomFee,
	).Scan(&res.ID)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
