package configs

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Глобальная переменная для базы данных
var DB *gorm.DB

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// Загружаем конфигурацию базы данных из переменных окружения
func LoadDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Host:     "138.124.81.155", // Указываем IP адрес сервера
		Port:     "5432",           // Указываем порт
		User:     "iam_user",       // Имя пользователя
		Password: "Agr3k0t_iam",    // Пароль
		DBName:   "iam_db",         // Имя базы данных
	}
}

// Функция для подключения к базе данных с использованием GORM
func ConnectDB() {
	config := LoadDatabaseConfig()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBName)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Логи запросов
	})

	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}

	log.Println("Подключение к базе данных успешно установлено!")
}
