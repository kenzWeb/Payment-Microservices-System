package client

import (
	"context"
	"github.com/user/payment-microservices/api/proto/order"
	"github.com/user/payment-microservices/pkg/grpcutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type OrderClient struct {
	client order.OrderServiceClient
}

func NewOrderClient(addr string) (*OrderClient, error) {
	conn, err := grpc.NewClient(
		addr, 
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(grpcutil.UnaryClientInterceptor()),
	)
	if err != nil { return nil, err }
	return &OrderClient{client: order.NewOrderServiceClient(conn)}, nil
}

func (c *OrderClient) CreateOrder(ctx context.Context, userID string, items []*order.OrderItem) (string, error) {
	resp, err := c.client.CreateOrder(ctx, &order.CreateOrderRequest{
		UserId: userID,
		Items:  items,
	})
	if err != nil { return "", err }
	return resp.OrderId, nil
}
