package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mulesoft.org/salsify-product-api/routers"
	"mulesoft.org/salsify-product-api/services"
)

func main() {
	router := gin.Default()

	// Health check route
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"api": "Composable Commerce Product API", "version": "1.0.10"})
	})

	// Product Routes
	pr := router.Group("/api/")
	routers.ProductRoutes(pr)

	services.ConnectDatabase()

	router.Run("0.0.0.0:8080")
}
