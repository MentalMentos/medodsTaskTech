package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"medodsTechTask/internal/data/request"
	"medodsTechTask/internal/data/response"
	"medodsTechTask/internal/model"
	"medodsTechTask/internal/repository"
	"medodsTechTask/pkg/logger"
	"medodsTechTask/pkg/utils"
)

type AuthService struct {
	repo   repository.Repository
	logger logger.Logger
}

func NewAuthService(repo repository.Repository, logger logger.Logger) *AuthService {
	return &AuthService{
		repo:   repo,
		logger: logger,
	}
}

func (s *AuthService) Login(ctx context.Context, req request.LoginRequest) (*response.TokenResponse, error) {
	// допустим пользователь уже прошёл проверку email/пароля
	userID := uuid.New()
	userRole := "user"
	ip := req.IP

	accessToken, jti, err := utils.GenerateJWT(userID, userRole, ip)
	if err != nil {
		s.logger.Fatal("[ LOGIN ]", " Failed to generate access token")
		return nil, err
	}

	rawRefresh, hashedRefresh, err := utils.GenerateSecureRefreshToken()
	if err != nil {
		s.logger.Fatal("[ LOGIN ] ", "Failed to generate refresh token")
		return nil, err
	}

	refresh := model.RefreshToken{
		UserID:      userID,
		AccessJTI:   jti,
		HashedToken: hashedRefresh,
		IPAddress:   ip,
		Used:        false,
	}

	if err := s.repo.SaveRefreshToken(ctx, refresh); err != nil {
		s.logger.Fatal("[ LOGIN ] ", "Failed to save refresh token")
		return nil, err
	}

	return &response.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: rawRefresh,
	}, nil
}

func (s *AuthService) RefreshToken(ctx context.Context, req request.UpdateTokenRequest) (*response.TokenResponse, error) {
	accessClaims, err := utils.ValidateJWT(req.AccessToken)
	if err != nil {
		s.logger.Debug("[ REFRESH ]", "Invalid access token")
		return nil, errors.New("invalid access token")
	}

	storedToken, err := s.repo.GetRefreshTokenByJTI(ctx, accessClaims.ID)
	if err != nil {
		s.logger.Debug("[ REFRESH ]", " Refresh token not found")
		return nil, errors.New("refresh token not found")
	}

	if bcrypt.CompareHashAndPassword([]byte(storedToken.HashedToken), []byte(req.RefreshToken)) != nil {
		s.logger.Debug("[ REFRESH ]", " Refresh token mismatch")
		return nil, errors.New("invalid refresh token")
	}

	if storedToken.Used {
		s.logger.Debug("[ REFRESH ]", " Refresh token already used")
		return nil, errors.New("refresh token already used")
	}

	if storedToken.IPAddress != req.IP {
		s.logger.Debug("[ REFRESH ]", " IP address mismatch — warning email could be sent")
	}
	if err := s.repo.MarkRefreshTokenUsed(ctx, storedToken.ID); err != nil {
		s.logger.Debug("[ REFRESH ]", " Failed to mark token as used")
	}

	newAccess, newJTI, err := utils.GenerateJWT(storedToken.UserID, accessClaims.Role, req.IP)
	if err != nil {
		return nil, err
	}

	newRawRefresh, newHashed, err := utils.GenerateSecureRefreshToken()
	if err != nil {
		return nil, err
	}

	newToken := model.RefreshToken{
		UserID:      storedToken.UserID,
		AccessJTI:   newJTI,
		HashedToken: newHashed,
		IPAddress:   req.IP,
		Used:        false,
	}

	if err := s.repo.SaveRefreshToken(ctx, newToken); err != nil {
		return nil, err
	}

	return &response.TokenResponse{
		AccessToken:  newAccess,
		RefreshToken: newRawRefresh,
	}, nil
}
