package service

import (
	"context"
	"github.com/MentalMentos/medodsTaskTech.git/internal/data/request"
	"github.com/MentalMentos/medodsTaskTech.git/internal/data/response"
	"github.com/MentalMentos/medodsTaskTech.git/internal/repository"
	"github.com/MentalMentos/medodsTaskTech.git/pkg/logger"
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
