package repository

import (
	"context"
	"gorm.io/gorm"
	"medodsTechTask/internal/model"
	"medodsTechTask/pkg/logger"
)

type RepoImpl struct {
	DB     *gorm.DB
	logger logger.Logger
}

func NewRepo(db *gorm.DB, logger logger.Logger) *RepoImpl {
	return &RepoImpl{
		db,
		logger,
	}
}

func (r *RepoImpl) SaveRefreshToken(ctx context.Context, token model.RefreshToken) error {
	return r.DB.WithContext(ctx).Create(&token).Error
}

func (r *RepoImpl) GetRefreshTokenByJTI(ctx context.Context, jti string) (model.RefreshToken, error) {
	var token model.RefreshToken
	err := r.DB.WithContext(ctx).
		Where("access_jti = ? AND used = FALSE", jti).
		First(&token).Error
	r.logger.Info("Getrefreshtokenbyjti", "successful")
	return token, err
}

func (r *RepoImpl) MarkRefreshTokenUsed(ctx context.Context, id int) error {
	err := r.DB.WithContext(ctx).
		Model(&model.RefreshToken{}).
		Where("id = ?", id).
		Update("used", true).Error
	if err != nil {
		r.logger.Debug("MarkRefreshTokenUsed", err.Error())
	}
	r.logger.Info("MarkRefreshTokenUsed", "successful")
	return nil
}
