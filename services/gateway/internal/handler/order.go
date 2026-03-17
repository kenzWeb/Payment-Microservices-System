package handler

import (
	"encoding/json"
	"net/http"
	"github.com/user/payment-microservices/services/gateway/internal/client"
	"github.com/user/payment-microservices/api/proto/order"
)

type OrderHandler struct {
	client *client.OrderClient
}

func NewOrderHandler(cl *client.OrderClient) *OrderHandler {
	return &OrderHandler{client: cl}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req struct {
		UserID string `json:"user_id"`
		Items  []struct {
			ProductID string  `json:"product_id"`
			Quantity  int32   `json:"quantity"`
			Price     float64 `json:"price"`
		} `json:"items"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	items := make([]*order.OrderItem, len(req.Items))
	for i, it := range req.Items {
		items[i] = &order.OrderItem{ProductId: it.ProductID, Quantity: it.Quantity, Price: it.Price}
	}

	id, err := h.client.CreateOrder(r.Context(), req.UserID, items)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"order_id": id})
}
