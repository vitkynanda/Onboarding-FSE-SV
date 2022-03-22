package routes

import (
	"go-api/config"
	"go-api/middleware"
	"go-api/repository/product_repository"
	"go-api/repository/role_repository"
	"go-api/repository/user_repository"
	"go-api/usecase/jwt_usecase"
	"go-api/usecase/product_usecase"
	"go-api/usecase/role_usecase"
	"go-api/usecase/user_usecase"

	"go-api/delivery/product_delivery"
	"go-api/delivery/role_delivery"
	"go-api/delivery/user_delivery"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)	

func HandlerRequest() {
	config.InitConfig()
	
	connection := config.Connect()
	roleRepository := role_repository.GetRoleRepository(connection)
	roleUsecase := role_usecase.GetRoleUsecase(roleRepository)
	roleDelivery := role_delivery.GetRoleDelivery(roleUsecase)
	productRepository := product_repository.GetProductRepository(connection)
	productUsecase := product_usecase.GetProductUsecase(productRepository)
	productDelivery := product_delivery.GetProductDelivery(productUsecase)
	userRepository := user_repository.GetUserRepository(connection)
	jwtAuth := jwt_usecase.GetJwtUsecase(userRepository)
	userUsecase := user_usecase.GetUserUsecase(userRepository, jwtAuth)
	userDelivery := user_delivery.GetUserDelivery(userUsecase)
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	  }))

	checkerRoutes := router.Group("/")
	checkerRoutes.Use(middleware.CheckerAuth(jwtAuth))

	{
		checkerRoutes.PUT("/products/:id/checked", productDelivery.UpdateProductData )
	}

	publisherRoutes := router.Group("/")
	publisherRoutes.Use(middleware.PublisherAuth(jwtAuth))

	{
		publisherRoutes.PUT("/products/:id/published", productDelivery.UpdateProductData )
	}

	adminRoutes := router.Group("/")
	adminRoutes.Use(middleware.AdminAuth(jwtAuth))
	{
		adminRoutes.PUT("/products/:id", productDelivery.UpdateProductData )
		adminRoutes.PUT("/users/:id", userDelivery.UpdateUserData )
		adminRoutes.DELETE("/products/:id", productDelivery.DeleteProductById )
		adminRoutes.DELETE("/users/:id", userDelivery.DeleteUserById )
	}

	protectedRoutes := router.Group("/")
	protectedRoutes.Use(middleware.JWTauth(jwtAuth))
	{
		router.POST("/products", productDelivery.CreateNewProduct )	
	}
	

	router.GET("/roles", roleDelivery.GetAllRole )	
	router.GET("/products", productDelivery.GetAllProducts )	
	router.GET("/products/:id", productDelivery.GetProductById )	
	
	router.POST("/login", userDelivery.UserLogin )	
	router.GET("/users", userDelivery.GetAllUsers )	
	router.GET("/users/:id", userDelivery.GetUserById )	
	router.POST("/users", userDelivery.CreateNewUser )	

	
	router.Run(":8001")
}