package client

import (
	"context"
	"time"
	"github.com/user/payment-microservices/api/proto/auth"
	"github.com/user/payment-microservices/pkg/grpcutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthClient struct {
	client  auth.AuthServiceClient
	breaker *grpcutil.CircuitBreaker
}

func NewAuthClient(addr string) (*AuthClient, error) {
	conn, err := grpc.NewClient(
		addr, 
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(grpcutil.UnaryClientInterceptor()),
	)
	if err != nil { return nil, err }
	return &AuthClient{
		client:  auth.NewAuthServiceClient(conn),
		breaker: grpcutil.NewBreaker(5, time.Minute),
	}, nil
}

func (c *AuthClient) GenerateToken(ctx context.Context, userID, email string) (string, error) {
	var token string
	err := c.breaker.Execute(func() error {
		resp, err := c.client.GenerateToken(ctx, &auth.GenerateTokenRequest{
			UserId: userID,
			Email:  email,
		})
		if err == nil { token = resp.Token }
		return err
	})
	return token, err
}
