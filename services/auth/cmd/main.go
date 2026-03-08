package main

import (
	"fmt"
	"net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"github.com/user/payment-microservices/pkg/config"
	"github.com/user/payment-microservices/pkg/logger"
	"github.com/user/payment-microservices/pkg/grpcutil"
	"github.com/user/payment-microservices/api/proto/auth"
	authCfg "github.com/user/payment-microservices/services/auth/internal/config"
	"github.com/user/payment-microservices/services/auth/internal/handler"
	"github.com/user/payment-microservices/services/auth/internal/service"
)

func main() {
	var cfg authCfg.Config
	config.Load(&cfg)
	log := logger.New(cfg.LogLevel)

	jwtSvc := service.NewJWTService(cfg.JWTSecret)
	h := handler.NewAuthHandler(jwtSvc)

	lis, _ := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPCPort))
	s := grpc.NewServer(grpc.UnaryInterceptor(grpcutil.UnaryServerInterceptor()))
	
	auth.RegisterAuthServiceServer(s, h)
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())

	log.Info("auth service starting", "port", cfg.GRPCPort)
	s.Serve(lis)
}
