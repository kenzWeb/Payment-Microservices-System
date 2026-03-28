package handler

import (
	"context"
	"github.com/user/payment-microservices/api/proto/auth"
	"github.com/user/payment-microservices/services/auth/internal/service"
)

type AuthHandler struct {
	auth.UnimplementedAuthServiceServer
	jwtService *service.JWTService
}

func NewAuthHandler(jwtService *service.JWTService) *AuthHandler {
	return &AuthHandler{jwtService: jwtService}
}

func (h *AuthHandler) GenerateToken(ctx context.Context, req *auth.GenerateTokenRequest) (*auth.GenerateTokenResponse, error) {
	token, err := h.jwtService.Generate(req.UserId, req.Email)
	if err != nil {
		return nil, err
	}
	return &auth.GenerateTokenResponse{Token: token}, nil
}

func (h *AuthHandler) ValidateToken(ctx context.Context, req *auth.ValidateTokenRequest) (*auth.ValidateTokenResponse, error) {
	userID, err := h.jwtService.Validate(req.Token)
	if err != nil {
		return &auth.ValidateTokenResponse{Valid: false}, nil
	}
	return &auth.ValidateTokenResponse{Valid: true, UserId: userID}, nil
}
