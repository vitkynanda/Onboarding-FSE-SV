package jwt_usecase

import (
	"go-api/repository/product_repository"
	"go-api/repository/user_repository"

	"github.com/golang-jwt/jwt"
)

type JwtUsecase interface {
	GenerateToken(string, string) (string, error)
	ValidateToken(string) (*jwt.Token, error)
	ValidateTokenAndGetUserId(string) (string, error)
	ValidateTokenAndGetRole(string)(string, string, error)
	CheckProductData(string, string) ( error)
}

type jwtUsecase struct {
	userRepo user_repository.UserRepository 
	productRepo product_repository.ProductRepository
}

func GetJwtUsecase(repo user_repository.UserRepository, product product_repository.ProductRepository) JwtUsecase {
	return &jwtUsecase{
		userRepo: repo,
		productRepo: product,
	}
}

