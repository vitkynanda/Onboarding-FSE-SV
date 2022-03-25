package product_usecase

import (
	"errors"
	"fmt"
	"go-api/helpers"
	"go-api/models/dto"
	"go-api/models/entity"

	"gorm.io/gorm"
)

func (product *productUsecase) GetAllProducts() dto.Response {
	productlist, err := product.productRepo.GetAllProducts()
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Data not found", err.Error(), 404)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", err.Error(), 500)
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
	productData, userData, err := product.productRepo.GetProductById(id)
	
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Data not found", err.Error(), 404)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", err.Error(), 500)
	}

	maker := dto.Action{}
	checker := dto.Action{}
	signer := dto.Action{}
 
	for _, user := range userData {
		if (user.ID == productData.MakerID) {
			maker.ID = user.ID
			maker.Name = user.Name
		} 
		if (user.ID == productData.CheckerID){
			checker.ID = user.ID
			checker.Name = user.Name
		} 
		if (user.ID == productData.SignerID){
			signer.ID = user.ID
			signer.Name = user.Name
		}
	}

	response := dto.ProductDetail{
		ID: productData.ID,
		Name: productData.Name,
		Description: productData.Description,
		Status: productData.Status,
		Maker: maker,
		Checker : checker,
		Signer: signer,
	}

	return helpers.ResponseSuccess("ok", nil, response, 200)
}

func (product *productUsecase) CreateNewProduct(newProduct dto.Product) dto.Response {
	fmt.Println(newProduct)
	productInsert := entity.Product{
		MakerID: newProduct.MakerID,
		Name: newProduct.Name,
		Description: newProduct.Description,
		Status: "inactive",
	}

	productData,  err := product.productRepo.CreateNewProduct(productInsert)
	
	 if err != nil {
		return helpers.ResponseError("Internal server error", err, 500)
	}

	return helpers.ResponseSuccess("ok", nil, map[string]interface{}{
		"id": productData.ID}, 201 )
}

func (product *productUsecase) UpdateProductData(productUpdate dto.Product, id string) dto.Response {
	productInsert := entity.Product{
		Name: productUpdate.Name,
		Description: productUpdate.Description,
	}

	_, err := product.productRepo.UpdateProductData(productInsert, id)
	 
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Data not found", err, 404)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", err, 500)
	}

	// if (productData.SignerID != "" || productData.CheckerID != "" ) {
	// 	return helpers.ResponseError("Forbidden access", errors.New("You cannot update this product because it is already checked or signed"), 403)
	// }

	return helpers.ResponseSuccess("ok", nil, map[string]interface{}{"id": id}, 200)
}

func (product *productUsecase) PublishedProduct(productUpdate dto.Product, id string) dto.Response {

fmt.Printf("%+v", productUpdate)
	productInsert := entity.Product{
		SignerID: productUpdate.SignerID,
		Name: productUpdate.Name,
		Description: productUpdate.Description,
	}
	_, err := product.productRepo.PublishedProduct(productInsert, id)
	 
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Data not found", err, 404)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", err, 500)
	}
	return helpers.ResponseSuccess("ok", nil, map[string]interface{}{"id": id}, 200)
}

func (product *productUsecase) CheckedProduct(productUpdate dto.Product, id string) dto.Response {
	productInsert := entity.Product{
		CheckerID: productUpdate.CheckerID,
		Name: productUpdate.Name,
		Description: productUpdate.Description,
	}
	_, err := product.productRepo.CheckedProduct(productInsert, id)
	 
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
