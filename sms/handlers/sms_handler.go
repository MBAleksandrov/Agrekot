package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sms/services" // Путь к сервисам, которые будут отправлять SMS и сообщения в Telegram
	"sms/utils"    // Путь к утилитам для обработки ответа

	"github.com/gin-gonic/gin"
)

// SendMessageHandler - обработчик для отправки сообщений через Telegram или SMS
func SendMessageHandler(c *gin.Context) {
	var req struct {
		Recipient string `json:"recipient"` // Получатель (номер телефона для SMS или chat_id для Telegram)
		Sender    string `json:"sender"`    // Имя отправителя
		Message   string `json:"message"`   // Сообщение
		Service   string `json:"service"`   // Тип сервиса: "sms" или "telegram"
	}

	// Проверка, что JSON валиден
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, "Invalid request", nil)
		return
	}

	// Определение сервиса для отправки сообщения
	if req.Service == "sms" {
		// Отправка через SMS (старый метод)
		err := services.SendSMS(req.Recipient, req.Sender, req.Message, "sms") // Добавили аргумент service
		if err != nil {
			utils.JSONResponse(c, http.StatusInternalServerError, "Failed to send SMS: "+err.Error(), nil)
			return
		}
		utils.JSONResponse(c, http.StatusOK, "SMS sent successfully", nil)
	} else if req.Service == "telegram" {
		// Отправка через Telegram
		// Получаем токен Telegram API из переменных окружения
		botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
		if botToken == "" {
			utils.JSONResponse(c, http.StatusInternalServerError, "Missing Telegram bot token", nil)
			return
		}

		// Отправляем сообщение в Telegram
		err := sendTelegramMessage(botToken, req.Recipient, req.Message)
		if err != nil {
			utils.JSONResponse(c, http.StatusInternalServerError, "Failed to send message to Telegram: "+err.Error(), nil)
			return
		}

		utils.JSONResponse(c, http.StatusOK, "Message sent to Telegram successfully", nil)
	} else {
		utils.JSONResponse(c, http.StatusBadRequest, "Invalid service type. Use 'sms' or 'telegram'", nil)
	}
}

// Функция для отправки сообщения через Telegram Bot API
func sendTelegramMessage(botToken, chatID, message string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)

	// Формируем данные для отправки
	data := map[string]string{
		"chat_id": chatID,
		"text":    message,
	}

	// Преобразуем данные в JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %v", err)
	}

	// Отправляем POST-запрос в Telegram API
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to send message to Telegram: %v", err)
	}
	defer resp.Body.Close()

	// Проверяем успешность ответа
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send message to Telegram, status code: %d", resp.StatusCode)
	}

	return nil
}
