package product_usecase

import (
	"go-api/models/dto"
	"go-api/repository/product_repository"
)

type ProductUsecase interface {
	GetAllProducts() dto.Response
	GetProductById(string) dto.Response
	CreateNewProduct(dto.Product) dto.Response
	UpdateProductData(dto.Product, string) dto.Response
	PublishedProduct(dto.Product, string) dto.Response
	CheckedProduct(dto.Product, string) dto.Response
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

type ProductUcaseTest struct {
	productRepo *product_repository.ProductRepositoryMock
}



