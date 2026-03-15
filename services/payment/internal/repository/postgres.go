package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type PaymentRepository struct {
	pool  *pgxpool.Pool
	redis *redis.Client
}

func NewPaymentRepository(pool *pgxpool.Pool, rds *redis.Client) *PaymentRepository {
	return &PaymentRepository{pool: pool, redis: rds}
}

func (r *PaymentRepository) IsProcessed(ctx context.Context, orderID string) (bool, error) {
	val, err := r.redis.Get(ctx, "payment:"+orderID).Result()
	if err == redis.Nil {
		return false, nil
	}
	return val == "processed", err
}

func (r *PaymentRepository) MarkProcessed(ctx context.Context, orderID string) error {
	return r.redis.Set(ctx, "payment:"+orderID, "processed", time.Hour*24).Err()
}

func (r *PaymentRepository) SavePayment(ctx context.Context, orderID string, amount float64, status string) error {
	_, err := r.pool.Exec(ctx, "INSERT INTO payments (order_id, amount, status) VALUES ($1, $2, $3)", orderID, amount, status)
	return err
}
