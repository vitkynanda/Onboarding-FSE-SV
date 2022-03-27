package product_usecase

import (
	"go-api/helpers"
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

type ProductUsecaseTest interface {
	GetAllProducts() dto.Response
	GetProductById(string) dto.Response
	UpdateProductData(string) dto.Response
	DeleteProductById(string) dto.Response
	CreateNewProduct(dto.Product) dto.Response
}

type productUsecaseTest struct{
	productRepo *product_repository.ProductRepositoryMock
}

func GetProductUsecaseTest(productRepositoryMock *product_repository.ProductRepositoryMock) ProductUsecaseTest {
	return &productUsecaseTest{
		productRepo: productRepositoryMock,
	}
}

func (usecase *productUsecaseTest) GetAllProducts() dto.Response {
	products, err := usecase.productRepo.GetAllProducts()
	if  products == nil {
		return helpers.ResponseError("Data not found", err, 404)
	} else {
		return helpers.ResponseSuccess("ok", nil, products, 200)
	}
}

func (usecase *productUsecaseTest) GetProductById(id string) dto.Response {
	product, err := usecase.productRepo.GetProductById(id)
	if  err != nil {
		return helpers.ResponseError("Data not found", err, 404)
	} else {
		return helpers.ResponseSuccess("ok", nil, product, 200)
	}
}
func (usecase *productUsecaseTest) UpdateProductData(id string) dto.Response {
	product, err := usecase.productRepo.UpdateProductData(id)
	
	if  err != nil {
		return helpers.ResponseError("Data not found", err, 404)
	} else {
		return helpers.ResponseSuccess("ok", nil, map[string]interface{}{"id": product.ID,}, 200)
	}
}

func (usecase *productUsecaseTest) DeleteProductById(id string) dto.Response {
	product, err := usecase.productRepo.DeleteProductById(id)
	
	if  err != nil {
		return helpers.ResponseError("Data not found", err, 404)
	} else {
		return helpers.ResponseSuccess("ok", nil, map[string]interface{}{"id": product.ID,}, 200)
	}
}
func (usecase *productUsecaseTest) CreateNewProduct(productInsert dto.Product) dto.Response {
	product, err := usecase.productRepo.CreateNewProduct(productInsert)
	if  err != nil {
		return helpers.ResponseError("Data not found", err, 404)
	} else {
		return helpers.ResponseSuccess("ok", nil, product, 200)
	}
}