package routes

import (
	"go-api/connection"
	"go-api/repository/user_repository"
	"go-api/usecase/user_usecase"

	"go-api/delivery/user_delivery"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)	

func HandlerRequest() {
	connection := connection.Connect()
	userRepository := user_repository.GetUserRepository(connection)
	userUsecase := user_usecase.GetUserUsecase(userRepository)
	userDelivery := user_delivery.GetUserDelivery(userUsecase)
	router := gin.Default()

	router.Use(cors.Default())
	router.GET("/users", userDelivery.GetAllUsers )	
	router.GET("/users/:id", userDelivery.GetUserById )	
	router.POST("/users", userDelivery.CreateNewUser )	
	router.PUT("/users/:id", userDelivery.UpdateUserData )
	router.DELETE("/users/:id", userDelivery.DeleteUserById )
	router.Run(":8001")
}