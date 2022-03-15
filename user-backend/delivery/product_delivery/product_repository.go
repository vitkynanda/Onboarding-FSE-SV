package product_delivery

import (
	"go-api/usecase/product_usecase"

	"github.com/gin-gonic/gin"
)

type ProductDelivery interface {
	GetAllProducts(*gin.Context)
	GetProductById(*gin.Context)
	CreateNewProduct(*gin.Context)
	UpdateProductData(*gin.Context)
	DeleteProductById(*gin.Context)
}

type productDelivery struct {
	productUsecase product_usecase.ProductUsecase
}

func GetProductDelivery(usecase product_usecase.ProductUsecase) ProductDelivery {
	return &productDelivery{
		productUsecase: usecase,
	}
}
