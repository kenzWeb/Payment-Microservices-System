package worker

import (
	"context"
	"time"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/user/payment-microservices/pkg/kafka"
)

type OutboxRelay struct {
	pool     *pgxpool.Pool
	producer *kafka.Producer
}

func NewOutboxRelay(pool *pgxpool.Pool, producer *kafka.Producer) *OutboxRelay {
	return &OutboxRelay{pool: pool, producer: producer}
}

func (r *OutboxRelay) Start(ctx context.Context) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done(): return
		case <-ticker.C:
			r.process(ctx)
		}
	}
}

func (r *OutboxRelay) process(ctx context.Context) {
	rows, _ := r.pool.Query(ctx, "SELECT id, topic, payload FROM outbox WHERE status = 'PENDING' LIMIT 10")
	defer rows.Close()

	for rows.Next() {
		var id, topic string
		var payload []byte
		rows.Scan(&id, &topic, &payload)

		if err := r.producer.Publish(ctx, topic, []byte(id), payload); err == nil {
			r.pool.Exec(ctx, "UPDATE outbox SET status = 'PROCESSED' WHERE id = $1", id)
		}
	}
}
