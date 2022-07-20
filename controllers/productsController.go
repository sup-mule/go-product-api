package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"mulesoft.org/salsify-product-api/models"
	"mulesoft.org/salsify-product-api/services"
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
		prodId, _ := strconv.ParseInt(c.Params.ByName("productId"), 0, 32)
		if newProd, err := services.SelectProductByID(int32(prodId)); err != nil {
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
		newID := services.InsertNewProduct(&newProduct)
		c.Header("Location", fmt.Sprintf("/api/products/%v", newID))
		c.Status(http.StatusCreated)
	}
}

func PutProductHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		productToUpdate, shouldReturn := validatePayload(c)
		if shouldReturn {
			return
		}
		prodId, _ := strconv.ParseInt(c.Param("productId"), 0, 32)
		if err := services.UpdateProduct(&productToUpdate, int32(prodId)); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusOK)
	}
}
