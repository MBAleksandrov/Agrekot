package main

import (
	"log"
	"pim/config"
	"pim/models"
	"pim/routes"
)

func main() {
	// Инициализация базы данных
	config.InitDatabase()

	// Автосоздание таблицы
	if err := config.DB.AutoMigrate(&models.Product{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Заполнение тестовыми данными
	seedDatabase()

	// Настройка маршрутов
	r := routes.SetupRoutes()

	// Запуск сервера
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func seedDatabase() {
	product := models.Product{
		Name:        "Томат Краснодарский",
		Description: "Свежий томат из Краснодарского края",
		Price:       100.50,
		Category:    "Овощи",
	}

	if err := config.DB.Create(&product).Error; err != nil {
		log.Printf("Failed to seed database: %v", err)
	} else {
		log.Println("Test data seeded successfully!")
	}
}
