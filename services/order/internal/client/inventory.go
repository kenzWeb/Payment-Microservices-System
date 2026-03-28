package client

import (
	"context"
	"time"
	"github.com/user/payment-microservices/api/proto/inventory"
	"github.com/user/payment-microservices/pkg/grpcutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type InventoryClient struct {
	client  inventory.InventoryServiceClient
	breaker *grpcutil.CircuitBreaker
}

func NewInventoryClient(addr string) (*InventoryClient, error) {
	conn, err := grpc.NewClient(
		addr, 
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(grpcutil.UnaryClientInterceptor()),
	)
	if err != nil { return nil, err }
	return &InventoryClient{
		client:  inventory.NewInventoryServiceClient(conn),
		breaker: grpcutil.NewBreaker(3, time.Second*30),
	}, nil
}

func (c *InventoryClient) CheckAndDeduct(ctx context.Context, productID string, quantity int32) (bool, error) {
	var success bool
	err := c.breaker.Execute(func() error {
		resp, err := c.client.DeductStock(ctx, &inventory.DeductStockRequest{
			ProductId: productID,
			Quantity:  quantity,
		})
		if err == nil { success = resp.Success }
		return err
	})
	return success, err
}
