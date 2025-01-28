package main

import (
	"iam/configs"
	"iam/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Подключение к базе данных
	configs.ConnectDatabase()

	// Настройка роутера
	r := gin.Default()

	// Подключение маршрутов аутентификации
	routes.AuthRoutes(r)

	// Запуск сервера
	r.Run(":8080")
}
