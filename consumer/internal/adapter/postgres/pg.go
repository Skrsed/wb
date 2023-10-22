package pgRepository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

/**
 * DB is a wrapper for PostgreSQL database connection
 * that uses pgxpool as database driver
 */
type DB struct {
	*pgxpool.Pool
}

type Credentials struct {
	User     string
	Password string
	Host     string
	Port     string
	DB       string
}

// NewDB creates a new PostgreSQL database instance
func NewDBConnection(ctx context.Context, cr *Credentials) (*DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		cr.User,
		cr.Password,
		cr.Host,
		cr.Port,
		cr.DB,
	)

	db, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return &DB{
		db,
	}, nil
}

// Close closes the database connection
func (db *DB) Close() {
	db.Pool.Close()
}
