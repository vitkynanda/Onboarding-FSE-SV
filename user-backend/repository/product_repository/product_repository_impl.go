package product_repository

import (
	"fmt"
	"go-api/models/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (repo *productRepository) GetAllProducts() ([]entity.Product, error) {
	products := []entity.Product{}
	err := repo.mysqlConnection.Model(&entity.Product{}).Scan(&products).Error
	if err != nil {
		return nil, err
	}

	if  len(products) <= 0 {
		return nil, gorm.ErrRecordNotFound
	}
	
	return products, nil
}

func (repo *productRepository) GetProductById(id string) (*entity.ProductDetail, error) {
	productDetail := entity.ProductDetail{}
	product := entity.Product{}
	users := []entity.User{}

	if err := repo.mysqlConnection.Where("id = ?", id).Find(&product).Error; err != nil {
		return nil, err
	}

	if (entity.Product{}) == product {
		return nil, gorm.ErrRecordNotFound
	}

	productDetail.ID = product.ID
	productDetail.Name = product.Name
	productDetail.Description = product.Description
	productDetail.Status = product.Status 

	err := repo.mysqlConnection.Where("name IN ?", []string{product.MakerID, product.CheckerID, product.MakerID}).Find(&users).Error;
	if err != nil {
		return nil, err
	}

	for i, user := range users {
		if (i == 0) {
			productDetail.MakerID = user.ID
			productDetail.MakerName = user.Name
		} else if (i == 1){
			productDetail.CheckerID = user.ID
			productDetail.CheckerName = user.Name
		} else {
			productDetail.SignerID = user.ID
			productDetail.SignerName = user.Name
		}
	}
	
	return &productDetail, nil
}

func (repo *productRepository) CreateNewProduct(product entity.Product) (*entity.Product,  error){
	product.ID = uuid.New().String()
	product.MakerID = "system"
	product.CheckerID = ""
	product.SignerID = ""
	
	if err := repo.mysqlConnection.Create(&product).Error; err != nil {
		return  nil, err
	}
	return &product,  nil
}

func (repo *productRepository) UpdateProductData(product entity.Product, id string, actionType string) (*entity.Product, error){
	switch actionType {
	case "published" :
		if err := repo.mysqlConnection.Model(&product).Where("id = ?", id).Updates(map[string]interface{}{"name": product.Name, "description": product.Description,"signer_id": "system", "status":"active"}).Error; err != nil {
				return nil, err
			}
	case "checked" :
		if err := repo.mysqlConnection.Model(&product).Where("id = ?", id).Updates(map[string]interface{}{"name": product.Name, "description": product.Description,"checker_id": "system"}).Error; err != nil {
				return nil, err
			}
	default :
		if err := repo.mysqlConnection.Model(&product).Where("id = ?", id).Updates(map[string]interface{}{"name": product.Name, "description": product.Description}).Error; err != nil {
			return nil, err
		}			
	}

	return &product, nil
}

func (repo *productRepository) DeleteProductById(id string) error {
		sql := "DELETE FROM products"
		sql = fmt.Sprintf("%s WHERE id = '%s'", sql, id)
		if err := repo.mysqlConnection.Raw(sql).Scan(entity.Product{}).Error; err != nil  {
			return err
		}

	return nil
}