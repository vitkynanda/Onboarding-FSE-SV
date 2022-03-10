package user_usecase

import (
	"go-api/models/dto"
	"go-api/repository/user_repository"
)

type UserUsecase interface {
	GetAllUsers() (dto.Response)
	GetUserById(string) (dto.Response)
	CreateNewUser(dto.User) (dto.Response)
	UpdateUserData(dto.User, string) (dto.Response)
	DeleteUserById(string) (dto.Response)
}

type userUsecase struct {
	userRepo user_repository.UserRepository
}

func GetUserUsecase(userRepository user_repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: userRepository,
	}
}
