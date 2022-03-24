package product_delivery

import (
	"go-api/helpers"
	"go-api/models/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (product *productDelivery) GetAllProducts(c *gin.Context) {
	response := product.productUsecase.GetAllProducts()
	// fmt.Printf("%+v", response)
	if (response.Status != "ok") {
		c.JSON(response.StatusCode, response)
		return
	}
	c.JSON(response.StatusCode, response)
}

func (product *productDelivery) GetProductById(c *gin.Context)    {
	id := c.Param("id")
	response := product.productUsecase.GetProductById(id)

	if (response.StatusCode == http.StatusNotFound) {
		c.JSON(http.StatusOK, response)
		return
	}
	if (response.Status != "ok") {
		c.JSON(response.StatusCode, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

func (product *productDelivery) CreateNewProduct(c *gin.Context)  {
	userId, _ := c.Get("user_id")
	userMakerId, _ := userId.(string)	

	request := dto.Product{
		MakerID: 	userMakerId,
	}

	if errBind := c.ShouldBindJSON(&request); errBind != nil {
		errorRes := helpers.ResponseError("Bad Request", errBind.Error(), 400  )
		c.JSON(errorRes.StatusCode, errorRes)
		return
	}

	response := product.productUsecase.CreateNewProduct(request)

	if (response.Status != "ok") {
		c.JSON(response.StatusCode, response)
		return
	}

	c.JSON(response.StatusCode, response)	
}

func (product *productDelivery) UpdateProductData(c *gin.Context) {
	id := c.Param("id")
	request := dto.Product{}

	if err := c.ShouldBindJSON(&request); err != nil {
		errorRes := helpers.ResponseError("Bad Request", err.Error(), 400)
		c.JSON(errorRes.StatusCode, errorRes)
		return
	}

	response := product.productUsecase.UpdateProductData(request, id)
	if (response.Status != "ok") {
		c.JSON(response.StatusCode, response)
		return
	}
	c.JSON(response.StatusCode, response)
}

func (product *productDelivery) PublishedProduct(c *gin.Context) {
	id := c.Param("id")
	userId, _ := c.Get("user_id")
	userSignerId, _ := userId.(string)

	request := dto.Product{
		SignerID: userSignerId,
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		errorRes := helpers.ResponseError("Bad Request", err.Error(), 400)
		c.JSON(errorRes.StatusCode, errorRes)
		return
	}

	response := product.productUsecase.PublishedProduct(request, id)
	if (response.Status != "ok") {
		c.JSON(response.StatusCode, response)
		return
	}

	c.JSON(response.StatusCode, response)
}

func (product *productDelivery) CheckedProduct(c *gin.Context) {
	id := c.Param("id")
	userId, _ := c.Get("user_id")
	userCheckerId, _ := userId.(string)
	
	request := dto.Product{
		CheckerID: userCheckerId,
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		errorRes := helpers.ResponseError("Bad Request", err.Error(), 400)
		c.JSON(errorRes.StatusCode, errorRes)
		return
	}

	response := product.productUsecase.CheckedProduct(request, id)
	if (response.Status != "ok") {
		c.JSON(response.StatusCode, response)
		return
	}
	
	c.JSON(response.StatusCode, response)
}

func (product *productDelivery) DeleteProductById(c *gin.Context) {
	id := c.Param("id")
	response := product.productUsecase.DeleteProductById(id)
	if (response.Status != "ok") {
		c.JSON(response.StatusCode, response)
		return
	}
	c.JSON(response.StatusCode, response)
}