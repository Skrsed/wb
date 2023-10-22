package pgRepository

/**
 * DeliveryRepository implements port.DeliveryRepository interface
 * and provides an access to the postgres database
 */
type PostgresRepository struct {
	db *DB
}

// NewDeliveryRepository creates a new Delivery pgRepository instance
func NewPostgresRepository(db *DB) *PostgresRepository {
	return &PostgresRepository{
		db,
	}
}
