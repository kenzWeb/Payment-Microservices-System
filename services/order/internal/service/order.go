package service

import (
	"context"
	"encoding/json"
	"github.com/user/payment-microservices/pkg/kafka"
	"github.com/user/payment-microservices/services/order/internal/client"
	"github.com/user/payment-microservices/services/order/internal/repository"
)

type OrderService struct {
	repo     *repository.OrderRepository
	inv      *client.InventoryClient
	producer *kafka.Producer
}

func NewOrderService(repo *repository.OrderRepository, inv *client.InventoryClient, producer *kafka.Producer) *OrderService {
	return &OrderService{repo: repo, inv: inv, producer: producer}
}

func (s *OrderService) CreateOrder(ctx context.Context, userID string, items []struct{ProductID string; Quantity int32; Price float64}) (string, error) {
	var total float64
	for _, item := range items {
		ok, err := s.inv.CheckAndDeduct(ctx, item.ProductID, item.Quantity)
		if err != nil || !ok { return "", err }
		total += item.Price * float64(item.Quantity)
	}

	payload, _ := json.Marshal(map[string]any{
		"user_id": userID,
		"amount": total,
	})

	return s.repo.CreateWithOutbox(ctx, userID, total, payload)
}
