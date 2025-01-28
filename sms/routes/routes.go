package routes

import (
	"sms/handlers" // Путь к вашим обработчикам

	"github.com/gin-gonic/gin"
)

// InitializeRoutes инициализирует все маршруты
func InitializeRoutes() {
	r := gin.Default()

	// Новый маршрут для отправки сообщений через SMS или Telegram
	r.POST("/send-message", handlers.SendMessageHandler) // Новый универсальный маршрут для отправки сообщений

	r.Run(":8080") // Запускаем сервер на порту 8080
}
