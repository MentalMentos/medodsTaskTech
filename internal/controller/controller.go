package controller

import (
	"github.com/gin-gonic/gin"
	"medodsTechTask/internal/data/request"
	"medodsTechTask/internal/service"
	"medodsTechTask/pkg/logger"
	"net/http"
)

type AuthController struct {
	authService service.Service
	logger      logger.Logger
}

func NewAuthController(authService *service.Service, logger logger.Logger) *AuthController {
	return &AuthController{
		authService: *authService,
		logger:      logger,
	}
}

func (controller *AuthController) Login(c *gin.Context) {
	var userRequest request.LoginRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		HandleError(c, &ApiError{
			Code:    http.StatusBadRequest,
			Message: "Invalid request payload",
		})
		return
	}

	userRequest.IP = GetClientIP(c)

	authResp, err := controller.authService.Login(c, userRequest)
	if err != nil {
		HandleError(c, err)
		return
	}

	JsonResponse(c, http.StatusOK, "Auth successful", authResp)
}

func (controller *AuthController) RefreshToken(c *gin.Context) {
	var tokenRequest request.UpdateTokenRequest
	if err := c.ShouldBindJSON(&tokenRequest); err != nil {
		HandleError(c, &ApiError{
			Code:    http.StatusBadRequest,
			Message: "Invalid request payload",
		})
		return
	}

	tokenRequest.IP = GetClientIP(c)

	authResp, err := controller.authService.RefreshToken(c, tokenRequest)
	if err != nil {
		HandleError(c, err)
		return
	}

	JsonResponse(c, http.StatusOK, "Token refreshed successful", authResp)
}
