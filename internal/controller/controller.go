package controller

import (
	"net/http"

	"github.com/MentalMentos/medodsTaskTech/internal/data/request"
	response2 "github.com/MentalMentos/medodsTaskTech/internal/data/response"
	"github.com/MentalMentos/medodsTaskTech/internal/service"
	"github.com/MentalMentos/medodsTaskTech/pkg/logger"
	"github.com/gin-gonic/gin"
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

func (controller *AuthController) Register(c *gin.Context) {
	var userRequest request.RegisterUserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		HandleError(c, &ApiError{Code: http.StatusBadRequest, Message: "Invalid request payload"})
		return
	}

	userRequest.IP = GetClientIP(c)
	authResp, err := controller.authService.Register(c, userRequest)
	if err != nil {
		HandleError(c, err)
		return
	}

	JsonResponse(c, http.StatusOK, "Registration successful", authResp)
}

func (controller *AuthController) Login(c *gin.Context) {
	var userRequest request.LoginRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		HandleError(c, &ApiError{Code: http.StatusBadRequest, Message: "Invalid request payload"})
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

func (controller *AuthController) UpdatePassword(c *gin.Context) {
	var userRequest request.UpdateUserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		HandleError(c, &ApiError{Code: http.StatusBadRequest, Message: "Invalid request payload"})
		return
	}

	userRequest.IP = GetClientIP(c)
	authResp, err := controller.authService.UpdatePassword(c, userRequest)
	if err != nil {
		HandleError(c, err)
		return
	}

	JsonResponse(c, http.StatusOK, "Password updated successful", authResp)
}

func (controller *AuthController) RefreshToken(c *gin.Context) {
	var userRequest request.UpdateTokenRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		HandleError(c, &ApiError{Code: http.StatusBadRequest, Message: "Invalid request payload"})
		return
	}

	authResp, err := controller.authService.GetAccessToken(c, userRequest.RefreshToken)
	if err != nil {
		HandleError(c, err)
		return
	}

	JsonResponse(c, http.StatusOK, "Token refreshed successful", authResp)
}

func (controller *AuthController) SendMessage(c *gin.Context) {
	var userRequest request.YaRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		HandleError(c, &ApiError{Code: http.StatusBadRequest, Message: "Invalid request payload"})
		return
	}

	if err := controller.authService.SendNotify(userRequest.ToSendNotificationReq()); err != nil {
		HandleError(c, err)
		return
	}

	JsonResponse(c, http.StatusOK, "Notify Done Successful", nil)
}

func (controller *AuthController) HandleYa(c *gin.Context) {
	var req request.Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Простой ответ
	response := response2.ResponseYa{
		Response: &response2.ResponseResponse{
			Text:       "Привет! Я навык для Алисы!",
			Tts:        "Привет! Я навык для Алисы!",
			EndSession: false,
		},
		Version: req.Version,
	}

	c.JSON(http.StatusOK, response)
}
