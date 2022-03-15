package product_usecase

import (
	"go-api/models/dto"
	"go-api/repository/product_repository"
)

type ProductUsecase interface {
	GetAllProducts() dto.Response
	GetProductById(string) dto.Response
	CreateNewProduct(dto.Product) dto.Response
	UpdateProductData(dto.Product, string, string) dto.Response
	DeleteProductById(string) dto.Response
}

type productUsecase struct{
	productRepo product_repository.ProductRepository
}

func GetProductUsecase(productRepository product_repository.ProductRepository) ProductUsecase {
	return &productUsecase{
		productRepo: productRepository,
	}
}
