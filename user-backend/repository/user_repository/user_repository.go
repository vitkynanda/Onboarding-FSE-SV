package user_repository

import (
	"go-api/models/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers() ([]entity.UserList, error)
	GetUserById(string) (*entity.UserDetail, error) 
	CreateNewUser(entity.User) (*entity.User, *entity.Role, error)
	UpdateUserData(entity.User, string) (*entity.User, error) 
	DeleteUserById( string) error
}

type userRepository struct {
	mysqlConnection *gorm.DB
}

func GetUserRepository(mysqlConn *gorm.DB) UserRepository  {
	return &userRepository{
		mysqlConnection: mysqlConn,
	}
}

