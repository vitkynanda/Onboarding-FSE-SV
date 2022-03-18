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
	request := dto.Product{}
	if err := c.ShouldBindJSON(&request); err != nil {
		errorRes := helpers.ResponseError("Bad Request", err.Error(), 400  )
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

	userSignerId, err := userId.(string)

	if err  {
		errRes := helpers.ResponseError("Forbidden Access", "access denied", 403)
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	request := dto.Product{
		CheckerID: userSignerId,
	}

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

func (product *productDelivery) CheckedProduct(c *gin.Context) {
	id := c.Param("id")
	userId, _ := c.Get("user_id")
	userCheckerId, err := userId.(string)
	
	if err  {
		errRes := helpers.ResponseError("Forbidden Access", "access denied", 403)
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	request := dto.Product{
		CheckerID: userCheckerId,
	}

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

func (product *productDelivery) DeleteProductById(c *gin.Context) {
	id := c.Param("id")
	response := product.productUsecase.DeleteProductById(id)
	if (response.Status != "ok") {
		c.JSON(response.StatusCode, response)
		return
	}
	c.JSON(response.StatusCode, response)
}