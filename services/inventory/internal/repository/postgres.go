package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type InventoryRepository struct {
	pool *pgxpool.Pool
}

func NewInventoryRepository(pool *pgxpool.Pool) *InventoryRepository {
	return &InventoryRepository{pool: pool}
}

func (r *InventoryRepository) GetStock(ctx context.Context, productID string) (int32, error) {
	var stock int32
	err := r.pool.QueryRow(ctx, "SELECT quantity FROM inventory WHERE product_id = $1", productID).Scan(&stock)
	if err != nil {
		return 0, fmt.Errorf("query row: %w", err)
	}
	return stock, nil
}

func (r *InventoryRepository) DeductStock(ctx context.Context, productID string, quantity int32) error {
	tag, err := r.pool.Exec(ctx, "UPDATE inventory SET quantity = quantity - $1 WHERE product_id = $2 AND quantity >= $1", quantity, productID)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("insufficient stock or product not found")
	}
	return nil
}
