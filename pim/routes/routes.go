package routes

import (
	"pim/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/products/:id", controllers.GetProductByID)

	return r
}
