package product_usecase

import (
	"errors"
	"go-api/helpers"
	"go-api/models/dto"
	"go-api/models/entity"

	"gorm.io/gorm"
)

func (product *productUsecase) GetAllProducts() dto.Response {
	userlist, err := product.productRepo.GetAllProducts()

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Data not found", err)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", err)
	}
	return helpers.ResponseSuccess("ok", nil, userlist)
}

func (product *productUsecase) GetProductById(id string) dto.Response {
	userData, err := product.productRepo.GetProductById(id)
	
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Data not found", err)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", err)
	}

	return helpers.ResponseSuccess("ok", nil, userData)
}

func (product *productUsecase) CreateNewProduct(newProduct dto.Product) dto.Response {
	userInsert := entity.Product{
		// ID: newUser.Id,
		Name: newProduct.Name,
		// Email: newUser.Email,
		// Personal_number: newUser.Personal_number,
		// Password: newUser.Password,
	}

	userData,  err := product.productRepo.CreateNewProduct(userInsert)
	
	 if err != nil {
		return helpers.ResponseError("Internal server error", err)
	}

	return helpers.ResponseSuccess("ok", nil, map[string]interface{}{
		"id": userData.ID} )
}

func (product *productUsecase) UpdateProductData(productUpdate dto.Product, id string) dto.Response {
	productInsert := entity.Product{
		// Name: userUpdate.Name,
		// Email: userUpdate.Email,
		// Personal_number: userUpdate.Personal_number,
		// Active: userUpdate.Active,
		// Password: userUpdate.Password,
		// RoleID: userUpdate.Role.Id,
	}
	_, err := product.productRepo.UpdateProductData(productInsert, id)
	 
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Data not found", err)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", err)
	}

	// userUpdate.Id = id
	return helpers.ResponseSuccess("ok", nil, map[string]interface{}{"id": id})
}

func (product *productUsecase) DeleteProductById(id string) dto.Response {
	
 err := product.productRepo.DeleteProductById(id)
 
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Data not found", err)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", err)
	}
	return helpers.ResponseSuccess("ok", nil, nil)
}
