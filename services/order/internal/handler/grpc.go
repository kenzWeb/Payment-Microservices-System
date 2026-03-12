package handler

import (
	"context"
	"github.com/user/payment-microservices/api/proto/order"
	"github.com/user/payment-microservices/services/order/internal/service"
)

type OrderHandler struct {
	order.UnimplementedOrderServiceServer
	svc *service.OrderService
}

func NewOrderHandler(svc *service.OrderService) *OrderHandler {
	return &OrderHandler{svc: svc}
}

func (h *OrderHandler) CreateOrder(ctx context.Context, req *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	items := make([]struct{ProductID string; Quantity int32; Price float64}, len(req.Items))
	for i, item := range req.Items {
		items[i] = struct{ProductID string; Quantity int32; Price float64}{item.ProductId, item.Quantity, item.Price}
	}

	id, err := h.svc.CreateOrder(ctx, req.UserId, items)
	if err != nil {
		return nil, err
	}

	return &order.CreateOrderResponse{OrderId: id, Status: "PENDING"}, nil
}
