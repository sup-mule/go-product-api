package services

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"mulesoft.org/salsify-product-api/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "products.db")

	if err != nil {
		panic("Failed to connect to the database!")
	}

	database.AutoMigrate(&models.Product{})

	// Load masterlist to dbase if empty
	var pd models.Product
	if err := database.First(&pd).Error; err != nil {
		for _, book := range models.ProductsMasterList {
			database.Create(&book)
		}
	}
	DB = database
}

func SelectAllProducts() []models.Product {
	var products []models.Product
	DB.Find(&products)
	return products
}

func SelectProductByID(id int64) (models.Product, error) {
	var product models.Product
	err := DB.First(&product, id).Error
	return product, err
}

func InsertNewProduct(product *models.Product) int64 {
	DB.Create(&product)
	return product.ID
}

func UpdateProduct(product *models.Product, productId int64) error {
	var curProduct models.Product
	if err := DB.First(&curProduct, productId).Error; err != nil {
		return err
	}
	DB.Model(&curProduct).Updates(product)
	return nil
}
