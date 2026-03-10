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
	"github.com/user/payment-microservices/pkg/grpcutil"
	"github.com/user/payment-microservices/api/proto/inventory"
	invCfg "github.com/user/payment-microservices/services/inventory/internal/config"
	"github.com/user/payment-microservices/services/inventory/internal/handler"
	"github.com/user/payment-microservices/services/inventory/internal/repository"
)

func main() {
	var cfg invCfg.Config
	config.Load(&cfg)
	log := logger.New(cfg.LogLevel)

	ctx := context.Background()
	pool, _ := postgres.NewPool(ctx, cfg.PostgresDSN)

	repo := repository.NewInventoryRepository(pool)
	h := handler.NewInventoryHandler(repo)

	s := grpc.NewServer(grpc.UnaryInterceptor(grpcutil.UnaryServerInterceptor()))
	inventory.RegisterInventoryServiceServer(s, h)
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())

	go func() {
		lis, _ := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPCPort))
		log.Info("inventory starting", "port", cfg.GRPCPort)
		s.Serve(lis)
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Info("shutting down inventory service...")
	s.GracefulStop()
	pool.Close()
}
