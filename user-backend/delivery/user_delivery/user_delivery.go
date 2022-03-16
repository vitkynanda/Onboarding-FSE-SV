package user_delivery

import (
	"go-api/usecase/user_usecase"

	"github.com/gin-gonic/gin"
)

type UserDelivery interface {
	UserLogin(*gin.Context) 
	GetAllUsers(*gin.Context)
	GetUserById(*gin.Context)
	CreateNewUser(*gin.Context)
	UpdateUserData(*gin.Context)
	DeleteUserById(*gin.Context)
}

type userDelivery struct {
	usecase user_usecase.UserUsecase
}

func GetUserDelivery(userUsecase user_usecase.UserUsecase) UserDelivery {
	return &userDelivery{
		usecase: userUsecase,
	}
}
