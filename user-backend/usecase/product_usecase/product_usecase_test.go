package product_usecase

import (
	"go-api/helpers"
	"go-api/models/dto"
	"go-api/models/entity"
	"go-api/repository/product_repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepository = &product_repository.ProductRepositoryMock{Mock : mock.Mock{}}
var productUse = GetProductUsecaseTest(productRepository)
func TestGetAllProduct(t *testing.T) {
	products := []entity.Product{}
	response := dto.Response{StatusCode: 200, Status: "ok", Error: nil, Data: products,}

	result := helpers.ResponseSuccess(response.Status, response.Error, response.Data,  response.StatusCode)
	productRepository.Mock.On("GetAllProducts").Return(products, nil)
	res := productUse.GetAllProducts()
	assert.Equal(t, result, res)
}

func TestGetProductById(t *testing.T){
	product := entity.Product{}
	response := dto.Response{StatusCode: 200, Status: "ok", Error: nil, Data: product,}

	result := helpers.ResponseSuccess(response.Status, response.Error, response.Data,  response.StatusCode)
	productRepository.Mock.On("GetProductById", "1").Return(product, nil)
	res := productUse.GetProductById("1")
	assert.Equal(t, result, res)
	assert.Nil(t, res.Error)
	assert.NotNil(t, res.Data)
}
func TestUpdateProductData(t *testing.T){
	product := entity.Product{ID : "1"}
	response := dto.Response{StatusCode: 200, Status: "ok", Error: nil, Data: map[string]interface{}{"id": product.ID},}

	result := helpers.ResponseSuccess(response.Status, response.Error, response.Data,  response.StatusCode)
	productRepository.Mock.On("UpdateProductData", "1").Return(product, nil)
	res := productUse.UpdateProductData("1")
	assert.Equal(t, result, res)
	assert.Nil(t, res.Error)
	assert.NotNil(t, res.Data)
}

func TestDeleteProductById(t *testing.T){
	product := entity.Product{ID: "1"}
	response := dto.Response{StatusCode: 200, Status: "ok", Error: nil, Data: map[string]interface{}{"id": product.ID},}

	result := helpers.ResponseSuccess(response.Status, response.Error, response.Data,  response.StatusCode)
	productRepository.Mock.On("DeleteProductById", "1").Return(product, nil)
	res := productUse.DeleteProductById("1")
	assert.Equal(t, result, res)
	assert.Nil(t, res.Error)
	assert.NotNil(t, res.Data)
}
func TestCreateNewProduct(t *testing.T){
	product := entity.Product{}
	response := dto.Response{StatusCode: 200, Status: "ok", Error: nil, Data: product,}

	result := helpers.ResponseSuccess(response.Status, response.Error, response.Data,  response.StatusCode)
	productRepository.Mock.On("CreateNewProduct").Return(product, nil)
	res := productUse.CreateNewProduct(dto.Product{ID: "1", Name: "Product 1",  Description: "Product 1",})
	assert.Equal(t, result, res)
	assert.Nil(t, res.Error)
	assert.NotNil(t, res.Data)
}

