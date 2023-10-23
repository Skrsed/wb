package pgRepository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type Querier interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

/**
 * DeliveryRepository implements port.DeliveryRepository interface
 * and provides an access to the postgres database
 */
type PostgresRepository struct {
	_db *DB
	_tx *pgx.Tx
	db  Querier
}

// NewDeliveryRepository creates a new Delivery pgRepository instance
func NewPostgresRepository(db *DB) *PostgresRepository {
	return &PostgresRepository{
		_db: db,
		_tx: nil,
		db:  db,
	}
}

func (pr *PostgresRepository) StartTx(
	ctx context.Context,
	txOptions pgx.TxOptions,
) (*PostgresRepository, error) {
	tx, err := pr._db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{
		_db: pr._db,
		_tx: &tx,
		db:  tx,
	}, nil
}

func (pr *PostgresRepository) Rollback(ctx context.Context) error {
	return (*pr._tx).Rollback(ctx)
}

func (pr *PostgresRepository) Commit(ctx context.Context) error {
	return (*pr._tx).Commit(ctx)
}

// func (pr *PostgresRepository) ExitTx(ctx context.Context) (*PostgresRepository, error) {
// 	err := (*pr._tx).Rollback(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &PostgresRepository{
// 		_db: pr._db,
// 		_tx: nil,
// 		db:  pr._db,
// 	}, nil
// }
