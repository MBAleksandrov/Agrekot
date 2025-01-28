package main

import (
	"log"
	"net/http"
	"sms/routes" // Путь к вашим маршрутам
)

func main() {
	// Инициализируем маршруты
	routes.InitializeRoutes()

	// Запускаем сервер
	log.Fatal("Server error: ", http.ListenAndServe(":8080", nil))
}
