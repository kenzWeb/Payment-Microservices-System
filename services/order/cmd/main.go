package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"github.com/user/payment-microservices/pkg/config"
	"github.com/user/payment-microservices/pkg/logger"
	"github.com/user/payment-microservices/pkg/postgres"
	"github.com/user/payment-microservices/pkg/kafka"
	"github.com/user/payment-microservices/pkg/grpcutil"
	"github.com/user/payment-microservices/api/proto/order"
	ordCfg "github.com/user/payment-microservices/services/order/internal/config"
	"github.com/user/payment-microservices/services/order/internal/handler"
	"github.com/user/payment-microservices/services/order/internal/repository"
	"github.com/user/payment-microservices/services/order/internal/service"
	"github.com/user/payment-microservices/services/order/internal/client"
	"github.com/user/payment-microservices/services/order/internal/worker"
)

func main() {
	var cfg ordCfg.Config
	config.Load(&cfg)
	log := logger.New(cfg.LogLevel)

	pool, _ := postgres.NewPool(context.Background(), cfg.PostgresDSN)
	invCl, _ := client.NewInventoryClient(cfg.InventoryAddr)
	prod, _ := kafka.NewProducer(cfg.KafkaBrokers)

	repo := repository.NewOrderRepository(pool)
	svc := service.NewOrderService(repo, invCl, prod)
	
	relay := worker.NewOutboxRelay(pool, prod)
	ctx, cancel := context.WithCancel(context.Background())
	go relay.Start(ctx)

	s := grpc.NewServer(grpc.UnaryInterceptor(grpcutil.UnaryServerInterceptor()))
	order.RegisterOrderServiceServer(s, handler.NewOrderHandler(svc))
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())

	go func() {
		lis, _ := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPCPort))
		log.Info("order starting", "port", cfg.GRPCPort)
		s.Serve(lis)
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Info("shutting down order service...")
	cancel()
	s.GracefulStop()
	prod.Close()
	pool.Close()
}
