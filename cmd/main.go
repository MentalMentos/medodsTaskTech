package main

import (
	"github.com/gin-gonic/gin"
	"medodsTechTask/internal/config"
	"medodsTechTask/internal/controller"
	"medodsTechTask/internal/model"
	"medodsTechTask/internal/repository"
	"medodsTechTask/internal/service"
	zaplogger "medodsTechTask/pkg/logger/zap"
	"net/http"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome Home!")
	})
	router.GET("/ip", func(c *gin.Context) {
		clientIP := c.ClientIP()
		c.JSON(200, gin.H{"ip": clientIP})
	})
	log := zaplogger.New()
	db := config.DatabaseConnection(log)
	db.Table("refresh_tokens").AutoMigrate(&model.RefreshToken{})

	authRepository := repository.NewRepository(db, log)
	authService := service.New(authRepository, log)
	authController := controller.NewAuthController(authService, log)

	authRoutes := router.Group("/auth_v1")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/refresh", authController.RefreshToken)
	}

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Main", "Failed to start server")
	}
}
