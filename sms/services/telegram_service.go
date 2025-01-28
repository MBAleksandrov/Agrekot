package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// SendTelegramMessage отправляет сообщение через Telegram Bot API
func SendTelegramMessage(botToken, chatID, message string) error {
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
