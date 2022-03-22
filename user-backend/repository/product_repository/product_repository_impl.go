package product_repository

import (
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

func (repo *productRepository) GetProductById(id string) (*entity.Product, []entity.User, error) {
	
	product := entity.Product{}
	users := []entity.User{}

	if err := repo.mysqlConnection.Where("id = ?", id).Find(&product).Error; err != nil {
		return nil,nil, err
	}

	if (entity.Product{} == product)  {
		return nil, nil,  gorm.ErrRecordNotFound
	}

	err := repo.mysqlConnection.Where("id IN ?", []string{product.MakerID, product.CheckerID, product.SignerID}).Find(&users).Error;
	if err != nil {
		return nil, nil, err
	}
	
	return &product, users, nil
}

func (repo *productRepository) CreateNewProduct(product entity.Product, userId string) (*entity.Product,  error){
	product.ID = uuid.New().String()
	product.MakerID = userId //example id
	product.CheckerID = ""
	product.SignerID = ""
	
	if err := repo.mysqlConnection.Create(&product).Error; err != nil {
		return  nil, err
	}
	return &product,  nil
}

func (repo *productRepository) UpdateProductData(product entity.Product, id string) (*entity.Product, error){
	
		 result := repo.mysqlConnection.Model(&product).Where("id = ?", id).Updates(map[string]interface{}{"name": product.Name, "description": product.Description}); 
		if result.RowsAffected == 0 {
			return nil, gorm.ErrRecordNotFound
					
	}

	return &product, nil
}

func (repo *productRepository) PublishedProduct(product entity.Product, id string) (*entity.Product, error){
	
		 result := repo.mysqlConnection.Model(&product).Where("id = ?", id).Updates(map[string]interface{}{"name": product.Name, "description": product.Description}); 
		if result.RowsAffected == 0 {
			return nil, gorm.ErrRecordNotFound
					
	}

	return &product, nil
}

func (repo *productRepository) CheckedProduct(product entity.Product, id string) (*entity.Product, error){
	
		 result := repo.mysqlConnection.Model(&product).Where("id = ?", id).Updates(map[string]interface{}{"name": product.Name, "description": product.Description}); 
		if result.RowsAffected == 0 {
			return nil, gorm.ErrRecordNotFound
					
	}

	return &product, nil
}

func (repo *productRepository) DeleteProductById(id string) error {
	result := repo.mysqlConnection.Where("id = ?", id).Delete(&entity.Product{})
	
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	
	return nil
}