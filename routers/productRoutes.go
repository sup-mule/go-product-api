package routers

import (
	"github.com/gin-gonic/gin"
	"mulesoft.org/salsify-product-api/controllers"
)

func ProductRoutes(g *gin.RouterGroup) {
	// /products
	g.GET("/products", controllers.GetAllProductsHandler())
	g.POST("/products", controllers.PostProductHandler())

	// /products/:productId
	g.GET("/products/:productId", controllers.GetProductByIdHandler())
	g.PUT("/products/:productId", controllers.PutProductHandler())
}
