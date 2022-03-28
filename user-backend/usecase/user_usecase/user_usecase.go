package user_usecase

import (
	"go-api/models/dto"
	"go-api/repository/user_repository"
	"go-api/usecase/jwt_usecase"
)

type UserUsecase interface {
	UserLogin(dto.UserLogin) (dto.Response)
	GetAllUsers() (dto.Response)
	GetUserById(string) (dto.Response)
	CreateNewUser(dto.User) (dto.Response)
	UpdateUserData(dto.UserUpdate, string) (dto.Response)
	DeleteUserById(string) (dto.Response)
}

type userUsecase struct {
	userRepo user_repository.UserRepository
	jwtAuth jwt_usecase.JwtUsecase
}

type UserUsecaseTest struct {
	userRepo *user_repository.UserRepositoryMock
}

func GetUserUsecase(userRepository user_repository.UserRepository, jwtAuthentication jwt_usecase.JwtUsecase)  UserUsecase {
	return &userUsecase{
		userRepo: userRepository,
		jwtAuth: jwtAuthentication,
	}
}
