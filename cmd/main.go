package main

import (
	"net/http"

	"github.com/MentalMentos/medodsTaskTech/internal/config"
	"github.com/MentalMentos/medodsTaskTech/internal/controller"
	"github.com/MentalMentos/medodsTaskTech/internal/model"
	"github.com/MentalMentos/medodsTaskTech/internal/repository"
	"github.com/MentalMentos/medodsTaskTech/internal/service"
	zaplogger "github.com/MentalMentos/medodsTaskTech/pkg/logger/zap"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	_ "github.com/go-playground/validator/v10"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil) // Доверять всем прокси
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome Home!")
	})
	router.GET("/ip", func(c *gin.Context) {
		// Получаем IP клиента
		clientIP := c.ClientIP() // Автоматически извлекает IP с учётом заголовков X-Forwarded-For, X-Real-IP
		c.JSON(200, gin.H{"ip": clientIP})
	})
	log := zaplogger.New()
	db := config.DatabaseConnection(log)
	//validate := validator.New()
	db.Table("users").AutoMigrate(&model.User{})

	authRepository := repository.NewRepository(db, log)
	authService := service.New(authRepository, log)
	authController := controller.NewAuthController(authService, log)

	authRoutes := router.Group("/auth_v1")
	{
		authRoutes.POST("/register", authController.Register)             // Регистрация
		authRoutes.POST("/login", authController.Login)                   // Вход
		authRoutes.POST("/refresh", authController.RefreshToken)          // Обновление токена
		authRoutes.PUT("/update-password", authController.UpdatePassword) // Обновление пароля
	}

	notifyRoutes := router.Group("/notify")
	{
		notifyRoutes.Any("/send-notify-webhook", authController.SendMessage)
	}

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Main", "Failed to start server")
	}
}
