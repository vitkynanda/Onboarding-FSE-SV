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

func (repo *productRepository) GetProductById(id string) (*entity.Product, error) {
	product := entity.Product{}
	
	err := repo.mysqlConnection.Model(&entity.Product{}).Where("id = ?", id).Find(&product).Error
	
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (repo *productRepository) CreateNewProduct(product entity.Product) (*entity.Product,  error){
	product.ID = uuid.New().String()
	
	if err := repo.mysqlConnection.Create(&product).Error; err != nil {
		return  nil, err
	}
	return &product,  nil
}

func (repo *productRepository) UpdateProductData(user entity.Product, id string) (*entity.Product, error){
	
	if err := repo.mysqlConnection.Model(&user).Where("id = ?", id).Updates(map[string]interface{}{

	}).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *productRepository) DeleteProductById(id string) error {
		sql := "DELETE FROM users"
		sql = fmt.Sprintf("%s WHERE id = '%s'", sql, id)
		if err := repo.mysqlConnection.Raw(sql).Scan(entity.Product{}).Error; err != nil  {
			return err
		}

	return nil
}