package user_repository

import (
	"fmt"
	"go-api/models/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (repo *userRepository) GetAllUsers() ([]entity.User, error) {
	users := []entity.User{}

	err := repo.mysqlConnection.Find(&users).Error

	if err != nil {
		return nil, err
	}

	if  len(users) <= 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return users, nil
}

func (repo *userRepository) GetUserById(id string) (*entity.User, error) {
	user := entity.User{}
	
	err := repo.mysqlConnection.Where("id = ?", id).Find(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *userRepository) CreateNewUser(user entity.User) (*entity.User, error){
	id := uuid.New()
	user.Id = id

	if err := repo.mysqlConnection.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *userRepository) UpdateUserData(user entity.User, id string ) (*entity.User, error){
	
	if err := repo.mysqlConnection.Model(&user).Where("id = ?", id).Updates(map[string]interface{}{
		"name": user.Name,
		"password": user.Password,
		"roleId": user.RoleId,
		"active": user.Active,
		"email": user.Email,
		"personal_number": user.Personal_number,
	}).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *userRepository) DeleteUserById(id string) error {
	sql := "DELETE FROM users"
	
	sql = fmt.Sprintf("%s WHERE id = '%s'", sql, id)

	if err := repo.mysqlConnection.Raw(sql).Scan(entity.User{}).Error; err != nil  {
		
		return err
	}
	// if err := repo.mysqlConnection.Delete(&entity.User{}, id).Error; err != nil  {
	// 	return err
	// }
	return nil
}