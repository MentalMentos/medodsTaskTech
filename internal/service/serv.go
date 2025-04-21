package service

import (
	"context"
	"medodsTechTask/internal/data/request"
	"medodsTechTask/internal/data/response"
	"medodsTechTask/internal/repository"
	"medodsTechTask/pkg/logger"
)

type AService interface {
	Login(ctx context.Context, req request.LoginRequest) (*response.TokenResponse, error)
	RefreshToken(ctx context.Context, req request.UpdateTokenRequest) (*response.TokenResponse, error)
}
type Service struct {
	AService
}

func New(repo repository.Repository, logger logger.Logger) *Service {
	return &Service{
		NewAuthService(repo, logger),
	}
}
