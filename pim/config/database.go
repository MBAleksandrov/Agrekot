package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error

	// Настройки подключения
	dsn := "host=138.124.81.155 user=pim_user password=Agr3k0t_pim dbname=pim_db port=5432 sslmode=disable TimeZone=Europe/Moscow"

	// Подключение к базе данных
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database. DSN: %s, Error: %v", dsn, err)
	}

	fmt.Println("Database connected successfully!")
}
