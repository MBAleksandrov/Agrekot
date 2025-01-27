package controllers

import (
	"net/http"
	"pim/config"
	"pim/models"

	"github.com/gin-gonic/gin"
)

func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}
