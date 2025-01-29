package main

import (
	"iam/configs"
	"iam/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Подключение к базе данных
	configs.ConnectDB()

	// Настройка роутера
	r := gin.Default()

	r.Use(cors.Default())

	// Подключение маршрутов аутентификации
	routes.AuthRoutes(r)

	// Запуск сервера
	r.Run(":8080")
}
