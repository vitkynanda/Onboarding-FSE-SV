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
	Role string `json:"role"`
	UserID string `json:"user_id"`
}

func (jwtAuth *jwtUsecase) GenerateToken(userId string, roleId string) (string, error) {
	data, err := jwtAuth.userRepo.GetRoleByRoleId(roleId)
	if err != nil {
		return "User not found", err
	}

	claim := CustomClaim{
		UserID: userId,
		Role: data.Title,
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

func (jwtAuth *jwtUsecase) ValidateTokenAndGetRole(token string) (string, string, error) {

	validatedToken, err := jwtAuth.ValidateToken(token)
	if err != nil {
		return "","", err
	}

	claims, ok := validatedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "","", errors.New("failed to claim token")
	}

	userData, err := jwtAuth.userRepo.GetUserById(claims["user_id"].(string))

	if err != nil {
		return "","", err
	}

	role, err := jwtAuth.userRepo.GetRoleByRoleId(userData.RoleID)
	if err != nil {
		return "","", err
	}

	return claims["user_id"].(string), role.Title, nil
}
func (jwtAuth *jwtUsecase) CheckProductData(id string, actionType string) ( error) {

	productData, _, err := jwtAuth.productRepo.GetProductById(id)

	if err != nil {
		return  err
	}



	if (actionType == "signer" && productData.CheckerID == "") {
		return  errors.New("This product need to be checked before signing")
	}

	if (actionType == "signer" && productData.SignerID == "") {
		return  nil
	}

	if (productData.SignerID != "" && productData.CheckerID != "" ) {
		return  errors.New("This product has been checked and signed")
	}

	if (productData.CheckerID != "") {
		return  errors.New("This product has been checked")
	}

	return  nil
}



