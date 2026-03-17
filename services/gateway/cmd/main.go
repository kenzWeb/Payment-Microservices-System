package main

import (
	"context"
	"fmt"
	"net/http"
	"github.com/google/uuid"
	"github.com/user/payment-microservices/pkg/config"
	"github.com/user/payment-microservices/pkg/logger"
	gwCfg "github.com/user/payment-microservices/services/gateway/internal/config"
	"github.com/user/payment-microservices/services/gateway/internal/client"
	"github.com/user/payment-microservices/services/gateway/internal/handler"
)

func traceMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tid := r.Header.Get("X-Trace-ID")
		if tid == "" { tid = uuid.NewString() }
		ctx := context.WithValue(r.Context(), logger.TraceIDKey, tid)
		w.Header().Set("X-Trace-ID", tid)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func main() {
	var cfg gwCfg.Config
	config.Load(&cfg)
	log := logger.New(cfg.LogLevel)

	authCl, _ := client.NewAuthClient(cfg.AuthAddr)
	ordCl, _ := client.NewOrderClient(cfg.OrderAddr)

	authHdl := handler.NewAuthHandler(authCl)
	ordHdl := handler.NewOrderHandler(ordCl)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /v1/auth/login", authHdl.Login)
	mux.HandleFunc("POST /v1/orders", ordHdl.CreateOrder)

	log.Info("gateway starting", "port", cfg.HTTPPort)
	addr := fmt.Sprintf(":%d", cfg.HTTPPort)
	http.ListenAndServe(addr, traceMiddleware(mux))
}
