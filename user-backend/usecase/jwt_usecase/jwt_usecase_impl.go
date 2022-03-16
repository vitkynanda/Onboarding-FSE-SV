package jwt_usecase

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type CustomClaim struct {
	jwt.StandardClaims
	RoleID string `json:"role_id"`
	UserID string `json:"user_id"`
}

func (jwtAuth *jwtUsecase) GenerateToken(userId string) (string, error) {
	claim := CustomClaim{
		UserID: userId,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
			Issuer:    os.Getenv("APP_NAME"),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claim)
	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func (jwtAuth *jwtUsecase) ValidateToken(token string) (*jwt.Token, error) {

	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
}

func (jwtAuth *jwtUsecase) ValidateTokenAndGetUserId(token string) (string, error) {
	validatedToken, err := jwtAuth.ValidateToken(token)
	if err != nil {
		return "", err
	}

	claims, ok := validatedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("failed to claim token")
	}

	return claims["user_id"].(string), nil
}
