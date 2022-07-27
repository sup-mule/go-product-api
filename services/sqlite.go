package services

import (
	"os"
	"path/filepath"

	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	
	"mulesoft.org/salsify-product-api/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	dataDir := os.Getenv("DATA_DIR")
	dbPath := filepath.Join(dataDir, "products.db")

	database, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

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

func SelectProductByID(productId int32) (models.Product, error) {
	var product models.Product
	err := DB.First(&product, productId).Error
	return product, err
}

func InsertNewProduct(product *models.Product) string {
	DB.Create(&product)
	return product.ProductID
}

func UpdateProduct(product *models.Product, productId int32) error {
	var curProduct models.Product
	if err := DB.First(&curProduct, productId).Error; err != nil {
		return err
	}
	DB.Model(&curProduct).Updates(product)
	return nil
}
