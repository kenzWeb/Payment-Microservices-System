package worker

import (
	"context"
	"encoding/json"
	"github.com/user/payment-microservices/pkg/kafka"
	"github.com/user/payment-microservices/services/payment/internal/repository"
)

type PaymentWorker struct {
	consumer *kafka.Consumer
	producer *kafka.Producer
	repo     *repository.PaymentRepository
}

func NewPaymentWorker(c *kafka.Consumer, p *kafka.Producer, r *repository.PaymentRepository) *PaymentWorker {
	return &PaymentWorker{consumer: c, producer: p, repo: r}
}

func (w *PaymentWorker) Start(ctx context.Context) {
	for {
		records, _ := w.consumer.Poll(ctx)
		for _, r := range records {
			var msg struct{ OrderID string `json:"order_id"`; Amount float64 `json:"amount"` }
			json.Unmarshal(r.Value, &msg)

			if processed, _ := w.repo.IsProcessed(ctx, msg.OrderID); processed {
				continue
			}

			status := "SUCCESS" // Mocked payment
			w.repo.SavePayment(ctx, msg.OrderID, msg.Amount, status)
			w.repo.MarkProcessed(ctx, msg.OrderID)

			resp, _ := json.Marshal(map[string]string{"order_id": msg.OrderID, "status": status})
			w.producer.Publish(ctx, "payment_completed", []byte(msg.OrderID), resp)
		}
	}
}
