package product_usecase

import (
	"errors"
	"go-api/helpers"
	"go-api/models/dto"
	"go-api/models/entity"

	"gorm.io/gorm"
)

func (product *productUsecase) GetAllProducts() dto.Response {
	productlist, err := product.productRepo.GetAllProducts()
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Data not found", err, 404)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", err, 500)
	}

	response := []dto.ProductList{}
	for _,product := range productlist {
		resProduct  := dto.ProductList{
			ID : product.ID,
			Name: product.Name,
			Description: product.Description,
			Status: product.Status,
		}
		response = append(response, resProduct)
	}

	return helpers.ResponseSuccess("ok", nil, response, 200)
}

func (product *productUsecase) GetProductById(id string) dto.Response {
	userData, err := product.productRepo.GetProductById(id)
	
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Data not found", err, 404)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", err, 500)
	}
	maker := dto.Action{
		ID: userData.MakerID,
		Name: userData.MakerName,
	}
	checker := dto.Action{
		ID: userData.CheckerID,
		Name: userData.CheckerName,
	}
	signer := dto.Action{
		ID: userData.SignerID,
		Name: userData.SignerName,
	}
	response := dto.ProductDetail{
		ID: userData.ID,
		Name: userData.Name,
		Description: userData.Description,
		Status: userData.Status,
		Maker: maker,
		Checker : checker,
		Signer: signer,
	}

	return helpers.ResponseSuccess("ok", nil, response, 200)
}

func (product *productUsecase) CreateNewProduct(newProduct dto.Product) dto.Response {
	userInsert := entity.Product{
		ID: newProduct.ID,
		Name: newProduct.Name,
		Description: newProduct.Description,
		Status: "inactive",
	}

	userData,  err := product.productRepo.CreateNewProduct(userInsert)
	
	 if err != nil {
		return helpers.ResponseError("Internal server error", err, 500)
	}

	return helpers.ResponseSuccess("ok", nil, map[string]interface{}{
		"id": userData.ID}, 200 )
}

func (product *productUsecase) UpdateProductData(productUpdate dto.Product, id string, actionType string) dto.Response {
	productInsert := entity.Product{
		Name: productUpdate.Name,
		Description: productUpdate.Description,
	}
	_, err := product.productRepo.UpdateProductData(productInsert, id, actionType)
	 
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Data not found", err, 404)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", err, 500)
	}
	return helpers.ResponseSuccess("ok", nil, map[string]interface{}{"id": id}, 200)
}

func (product *productUsecase) DeleteProductById(id string) dto.Response {
	
 err := product.productRepo.DeleteProductById(id)
 
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Data not found", err, 404)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", err, 500)
	}
	return helpers.ResponseSuccess("ok", nil, nil, 200)
}
