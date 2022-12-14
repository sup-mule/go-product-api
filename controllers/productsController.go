package controllers

import (
	"fmt"
	"net/http"
	"sync/atomic"

	"github.com/gin-gonic/gin"
	"mulesoft.org/go-product-api/models"
	"mulesoft.org/go-product-api/services"
)

func validatePayload(c *gin.Context) (models.Product, bool) {
	var newProduct models.Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return models.Product{}, true
	}
	return newProduct, false
}

func GetAllProductsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, services.SelectAllProducts())
	}
}

func GetProductByIdHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		prodId := c.Params.ByName("productId")
		if newProd, err := services.SelectProductByID(prodId); err != nil {
			c.JSON(http.StatusNotFound, err.Error())
		} else {
			c.JSON(http.StatusOK, newProd)
		}
	}
}

func PostProductHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// validate json payload
		newProduct, shouldReturn := validatePayload(c)
		if shouldReturn {
			return
		}
		newID := atomic.AddUint32(&models.IdCounter, 1)
		newProduct.ProductID = fmt.Sprintf("PRD%03d", newID)
		services.InsertNewProduct(&newProduct)
		c.Header("Location", fmt.Sprintf("/api/products/%v", newProduct.ProductID))
		c.Status(http.StatusCreated)
	}
}

func PutProductHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		productToUpdate, shouldReturn := validatePayload(c)
		if shouldReturn {
			return
		}
		prodId := c.Param("productId")
		if err := services.UpdateProduct(&productToUpdate, prodId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusOK)
	}
}
