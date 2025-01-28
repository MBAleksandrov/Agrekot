package services

import (
	"fmt"
)

// SendSigmaSMS отправляет SMS через Sigma API (заглушка для тестирования)
func SendSMS(apiToken, recipient, sender, message string) error {
	// Заглушка: Логируем сообщение, но не отправляем реально
	fmt.Printf("Simulated SMS sent to %s from %s: %s\n", recipient, sender, message)

	// Вместо отправки SMS, просто возвращаем nil для успешной симуляции
	return nil
}
