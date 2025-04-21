package repository

import (
	"context"
	"gorm.io/gorm"
	"medodsTechTask/internal/model"
	"medodsTechTask/pkg/logger"
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
