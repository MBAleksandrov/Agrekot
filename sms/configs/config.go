package configs

import (
	"os"
)

type Config struct {
	DBHost     string // Хост базы данных
	DBPort     string // Порт базы данных
	DBUser     string // Пользователь базы данных
	DBPassword string // Пароль базы данных
	DBName     string // Имя базы данных
	SSLMode    string // Режим SSL (например, disable, require)
	SMSAPIKey  string // Ключ API для SMS-шлюза
}

// LoadConfig загружает конфигурацию из переменных окружения
func LoadConfig() Config {
	return Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "sms_user"),
		DBPassword: getEnv("DB_PASSWORD", "Agr3k0t_sms"),
		DBName:     getEnv("DB_NAME", "sms_db"),
		SSLMode:    getEnv("DB_SSLMODE", "disable"),
		SMSAPIKey:  getEnv("SMS_API_KEY", "test_api_key"),
	}
}

// getEnv возвращает значение переменной окружения или значение по умолчанию
func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
