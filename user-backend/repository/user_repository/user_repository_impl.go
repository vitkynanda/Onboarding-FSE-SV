package user_repository

import (
	"errors"
	"fmt"
	"go-api/helpers"
	"go-api/models/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (repo *userRepository) GetRoleByRoleId(id string) (*entity.Role, error) {
	role := entity.Role{}
	result := repo.mysqlConnection.Where("id = ?", id).Find(&role)
	if (result.RowsAffected == 0) {
		return nil, gorm.ErrRecordNotFound
	}

	return &role, nil
}

func (repo *userRepository) GetUserByPN(pn string) (*entity.User, error) {
	user := entity.User{}
	result := repo.mysqlConnection.Where("personal_number = ?", pn).Find(&user)
	if (result.RowsAffected == 0) {
		return nil, gorm.ErrRecordNotFound
	}

	return &user, nil
}

func (repo *userRepository) GetAllUsers() ([]entity.UserList, error) {
	users := []entity.UserList{}
	err := repo.mysqlConnection.Model(&entity.User{}).Select("users.name, users.active, users.id, roles.title, users.role_id").Joins("left join roles on roles.id = users.role_id").Scan(&users).Error
	if err != nil {
		return nil, err
	}

	if  len(users) <= 0 {
		return nil, gorm.ErrRecordNotFound
	}
	
	return users, nil
}

func (repo *userRepository) GetUserById(id string) (*entity.UserDetail, error) {
	user := entity.UserDetail{}
	err := repo.mysqlConnection.Model(&entity.User{}).Where("users.id = ?", id).Select("users.name, users.active, users.email, users.personal_number, users.id, roles.title, users.role_id").Joins("left join roles on roles.id = users.role_id").Scan(&user).Error
	if err != nil {
		return nil, err
	}

	if (user == entity.UserDetail{}) {
		return nil, gorm.ErrRecordNotFound
	} 

	return &user, nil
}

func (repo *userRepository) CreateNewUser(user entity.User) (*entity.User,  error){
	role := entity.Role{}
	
	result := repo.mysqlConnection.Where("personal_number = ?", user.Personal_number).Find(&user)
	if (result.RowsAffected > 0)  {
		return nil,  errors.New("Personal number already registered")
	}

	user.ID = uuid.New().String()
	hash,_ := helpers.HashPassword(user.Password)
	user.Password = hash
	
	if err := repo.mysqlConnection.Where("title = ?", "viewer").Find(&role).Error; err != nil {
		return nil,  err
	}

	user.RoleID = role.ID

	if err := repo.mysqlConnection.Create(&user).Error; err != nil {
		return nil,  err
	}

	return &user,  nil
}

func (repo *userRepository) UpdateUserData(user entity.User, id string) (*entity.User, error){

	checkPN :=  repo.mysqlConnection.Where("personal_number = ?", user.Personal_number).Find(&user)
	if checkPN.RowsAffected > 0 {
		return nil, errors.New("Personal number already taken")
	}

	result:= repo.mysqlConnection.Model(&user).Where("id = ?", id).Updates(&user)
	
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &user, nil
}

func (repo *userRepository) DeleteUserById(id string) error {
	sql := "DELETE FROM users"
	sql = fmt.Sprintf("%s WHERE id = '%s'", sql, id)
	result := repo.mysqlConnection.Raw(sql).Scan(entity.User{})
	if (result.RowsAffected == 0)  {
		return gorm.ErrRecordNotFound
	}
	// if err := repo.mysqlConnection.Delete(&entity.User{}, id).Error; err != nil  {
	// 	return err
	// }
	return nil
}