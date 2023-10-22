package pgRepository

import (
	"consumer/internal/core/domain"
	"context"
)

// CreateItem creates a new Item record in the database
// TODO: Try to answer why we're returning pointer insted of actual value? Where that value is? Leaking or not?
func (pr *PostgresRepository) CreateItems(ctx context.Context, Items *[]*domain.Item) (*[]*domain.Item, error) {
	sql := `INSERT INTO items (
		chrt_id,
		track_number,
		price,
		rid,
		name,
		sale,
		size,
		total_price,
		nm_id,
		brand,
		status,
		order_uid
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	RETURNING id` // test perfomance

	res := make([]*domain.Item, len(*Items))

	// TODO: OPTIMIZE THAT SHIT!!!!!
	for i, Item := range *Items {
		itemRes := Item
		err := pr.db.QueryRow(ctx, sql,
			&Item.ChrtId,
			&Item.TrackNumber,
			&Item.Price,
			&Item.Rid,
			&Item.Name,
			&Item.Sale,
			&Item.Size,
			&Item.TotalPrice,
			&Item.NmId,
			&Item.Brand,
			&Item.Status,
		).Scan(&itemRes.ID)

		res[i] = itemRes

		if err != nil {
			return nil, err
		}
	}

	return &res, nil
}

func (pr *PostgresRepository) GetItemsByOrderUid(
	ctx context.Context,
	uid string,
) (*[]*domain.Item, error) {
	// TODO: extend with fields
	sql := `SELECT * FROM items WHERE uid = $1`

	rows, err := pr.db.Query(ctx, sql, uid)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*domain.Item

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var item domain.Item
		err := rows.Scan(
			&item.ID,
			&item.ChrtId,
			&item.TrackNumber,
			&item.Price,
			&item.Rid,
			&item.Name,
			&item.Sale,
			&item.Size,
			&item.TotalPrice,
			&item.NmId,
			&item.Brand,
			&item.Status,
			&item.OrderUid,
		)

		if err != nil {
			return nil, err
		}

		items = append(items, &item)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &items, nil
}
