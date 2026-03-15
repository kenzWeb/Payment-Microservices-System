package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"github.com/redis/go-redis/v9"
	"github.com/user/payment-microservices/pkg/config"
	"github.com/user/payment-microservices/pkg/logger"
	"github.com/user/payment-microservices/pkg/postgres"
	"github.com/user/payment-microservices/pkg/kafka"
	payCfg "github.com/user/payment-microservices/services/payment/internal/config"
	"github.com/user/payment-microservices/services/payment/internal/repository"
	"github.com/user/payment-microservices/services/payment/internal/worker"
)

func main() {
	var cfg payCfg.Config
	config.Load(&cfg)
	log := logger.New(cfg.LogLevel)
	ctx, cancel := context.WithCancel(context.Background())

	pool, _ := postgres.NewPool(ctx, cfg.PostgresDSN)
	rds := redis.NewClient(&redis.Options{Addr: cfg.RedisAddr})
	cons, _ := kafka.NewConsumer(cfg.KafkaBrokers, "payment-group", []string{"order_created"})
	prod, _ := kafka.NewProducer(cfg.KafkaBrokers)

	repo := repository.NewPaymentRepository(pool, rds)
	w := worker.NewPaymentWorker(cons, prod, repo)

	go func() {
		log.Info("payment service starting consumer")
		w.Start(ctx)
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Info("shutting down payment service...")
	cancel()
	cons.Close()
	prod.Close()
	pool.Close()
}
