package routes

import (
	"go-api/connection"
	"go-api/repository/product_repository"
	"go-api/repository/user_repository"
	"go-api/usecase/product_usecase"
	"go-api/usecase/user_usecase"

	"go-api/delivery/product_delivery"
	"go-api/delivery/user_delivery"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)	

func HandlerRequest() {
	connection := connection.Connect()
	productRepository := product_repository.GetProductRepository(connection)
	productUsecase := product_usecase.GetProductUsecase(productRepository)
	productDelivery := product_delivery.GetProductDelivery(productUsecase)
	userRepository := user_repository.GetUserRepository(connection)
	userUsecase := user_usecase.GetUserUsecase(userRepository)
	userDelivery := user_delivery.GetUserDelivery(userUsecase)
	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/products", productDelivery.GetAllProducts )	
	router.GET("/products/:id", productDelivery.GetProductById )	
	router.POST("/products", productDelivery.CreateNewProduct )	
	router.PUT("/products/:id", productDelivery.UpdateProductData )
	router.PUT("/products/:id/:type/", productDelivery.UpdateProductData )
	router.DELETE("/products/:id", productDelivery.DeleteProductById )

	router.POST("/login", userDelivery.UserLogin )	
	router.GET("/users", userDelivery.GetAllUsers )	
	router.GET("/users/:id", userDelivery.GetUserById )	
	router.POST("/users", userDelivery.CreateNewUser )	
	router.PUT("/users/:id", userDelivery.UpdateUserData )
	router.DELETE("/users/:id", userDelivery.DeleteUserById )
	router.Run(":8001")
}