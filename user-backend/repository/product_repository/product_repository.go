package product_repository

import (
	"go-api/models/entity"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAllProducts() ([]entity.Product, error)
	GetProductById(string) (*entity.Product, []entity.User, error)
	CreateNewProduct(entity.Product)(*entity.Product, error)
	UpdateProductData(entity.Product, string) (*entity.Product, error) 
	PublishedProduct(entity.Product, string) (*entity.Product, error) 
	CheckedProduct(entity.Product, string) (*entity.Product, error) 
	DeleteProductById(string) error
}

type productRepository struct{
	mysqlConnection *gorm.DB
}

func GetProductRepository(mysqlConn *gorm.DB) ProductRepository{
	return &productRepository{
		mysqlConnection: mysqlConn,
	}
}
