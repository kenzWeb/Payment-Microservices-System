package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"github.com/user/payment-microservices/pkg/config"
	"github.com/user/payment-microservices/pkg/logger"
	"github.com/user/payment-microservices/pkg/kafka"
	notCfg "github.com/user/payment-microservices/services/notification/internal/config"
	"github.com/user/payment-microservices/services/notification/internal/worker"
)

func main() {
	var cfg notCfg.Config
	config.Load(&cfg)
	log := logger.New(cfg.LogLevel)
	ctx, cancel := context.WithCancel(context.Background())

	cons, _ := kafka.NewConsumer(cfg.KafkaBrokers, "notif-group", []string{"order_created", "payment_completed"})
	w := worker.NewNotificationWorker(cons, log)

	go func() {
		log.Info("notification service starting consumer")
		w.Start(ctx)
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Info("shutting down notification service...")
	cancel()
	cons.Close()
}
