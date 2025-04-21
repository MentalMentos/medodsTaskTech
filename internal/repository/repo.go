package repository

import (
	"context"
	"github.com/MentalMentos/medodsTaskTech.git/internal/model"
	"github.com/MentalMentos/medodsTaskTech.git/pkg/logger"
	"gorm.io/gorm"
)

type Repository interface {
	SaveRefreshToken(ctx context.Context, token model.RefreshToken) error
	GetRefreshTokenByJTI(ctx context.Context, jti string) (model.RefreshToken, error)
	MarkRefreshTokenUsed(ctx context.Context, id int) error
}

type Repo struct {
	Repository
}

func NewRepository(db *gorm.DB, mylogger logger.Logger) *Repo {
	return &Repo{NewRepo(db, mylogger)}
}
