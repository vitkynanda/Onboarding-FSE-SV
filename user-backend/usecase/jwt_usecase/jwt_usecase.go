package jwt_usecase

import (
	"go-api/repository/user_repository"

	"github.com/golang-jwt/jwt"
)

type JwtUsecase interface {
	GenerateToken(string) (string, error)
	ValidateToken(string) (*jwt.Token, error)
	ValidateTokenAndGetUserId(string) (string, error)
}

type jwtUsecase struct {
	userRepo user_repository.UserRepository 
}

func GetJwtUsecase(repo user_repository.UserRepository) JwtUsecase {
	return &jwtUsecase{
		userRepo: repo,
	}
}

