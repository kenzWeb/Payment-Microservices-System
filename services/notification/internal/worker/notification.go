package worker

import (
	"context"
	"log/slog"
	"github.com/user/payment-microservices/pkg/kafka"
)

type NotificationWorker struct {
	consumer *kafka.Consumer
	log      *slog.Logger
}

func NewNotificationWorker(c *kafka.Consumer, log *slog.Logger) *NotificationWorker {
	return &NotificationWorker{consumer: c, log: log}
}

func (w *NotificationWorker) Start(ctx context.Context) {
	for {
		records, _ := w.consumer.Poll(ctx)
		for _, r := range records {
			w.log.Info("notification sent", "topic", r.Topic, "key", string(r.Key), "payload", string(r.Value))
		}
	}
}
