package handler

import (
	"context"
	"github.com/user/payment-microservices/api/proto/inventory"
	"github.com/user/payment-microservices/services/inventory/internal/repository"
)

type InventoryHandler struct {
	inventory.UnimplementedInventoryServiceServer
	repo *repository.InventoryRepository
}

func NewInventoryHandler(repo *repository.InventoryRepository) *InventoryHandler {
	return &InventoryHandler{repo: repo}
}

func (h *InventoryHandler) CheckStock(ctx context.Context, req *inventory.CheckStockRequest) (*inventory.CheckStockResponse, error) {
	stock, err := h.repo.GetStock(ctx, req.ProductId)
	if err != nil {
		return nil, err
	}
	return &inventory.CheckStockResponse{Available: stock >= req.Quantity}, nil
}

func (h *InventoryHandler) DeductStock(ctx context.Context, req *inventory.DeductStockRequest) (*inventory.DeductStockResponse, error) {
	err := h.repo.DeductStock(ctx, req.ProductId, req.Quantity)
	if err != nil {
		return &inventory.DeductStockResponse{Success: false}, nil
	}
	return &inventory.DeductStockResponse{Success: true}, nil
}
