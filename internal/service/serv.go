package service

import (
	"github.com/MentalMentos/medodsTaskTech/internal/repository"
	"github.com/MentalMentos/medodsTaskTech/pkg/logger"
)

type Service struct {
	*AuthService
}

func New(repo repository.Repository, logger logger.Logger) *Service {
	return &Service{
		NewAuthService(repo, logger),
	}
}
