package utils

import (
	"github.com/gin-gonic/gin"
)

// JSONResponse отправляет стандартный JSON-ответ
func JSONResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, gin.H{
		"message": message,
		"data":    data,
	})
}
