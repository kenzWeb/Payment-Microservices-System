package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepository struct {
	pool *pgxpool.Pool
}

func NewOrderRepository(pool *pgxpool.Pool) *OrderRepository {
	return &OrderRepository{pool: pool}
}

func (r *OrderRepository) CreateWithOutbox(ctx context.Context, userID string, total float64, payload []byte) (string, error) {
	tx, err := r.pool.Begin(ctx)
	if err != nil { return "", err }
	defer tx.Rollback(ctx)

	var id string
	err = tx.QueryRow(ctx, "INSERT INTO orders (user_id, total_amount, status) VALUES ($1, $2, 'PENDING') RETURNING id", userID, total).Scan(&id)
	if err != nil { return "", err }

	_, err = tx.Exec(ctx, "INSERT INTO outbox (topic, payload, status) VALUES ($1, $2, 'PENDING')", "order_created", payload)
	if err != nil { return "", err }

	return id, tx.Commit(ctx)
}

func (r *OrderRepository) UpdateStatus(ctx context.Context, id, status string) error {
	_, err := r.pool.Exec(ctx, "UPDATE orders SET status = $1 WHERE id = $2", status, id)
	return err
}
