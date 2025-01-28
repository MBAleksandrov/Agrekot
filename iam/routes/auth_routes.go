package routes

import (
	"iam/handlers"

	"github.com/gin-gonic/gin"
)

// AuthRoutes регистрирует маршруты для аутентификации
func AuthRoutes(router *gin.Engine) {
	// Группа маршрутов для /auth
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", handlers.Register)
		// Здесь можно добавить другие маршруты:
		// authGroup.POST("/login", handlers.Login)
		// authGroup.POST("/logout", handlers.Logout)
	}
}
