package user_delivery

import (
	"go-api/helpers"
	"go-api/models/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)



func (res *userDelivery) UserLogin(c *gin.Context) {
	request := dto.UserLogin{}
	if err := c.ShouldBindJSON(&request); err != nil {
		errorRes := helpers.ResponseError("Bad Request", err, 400)
		c.JSON(errorRes.StatusCode, errorRes)
		return
	}
	response := res.usecase.UserLogin(request)
	
	if (response.Status != "ok") {
		c.JSON(response.StatusCode, response)
		return
	}
	c.JSON(response.StatusCode, response)
}

func (res *userDelivery) GetAllUsers(c *gin.Context) {
	response := res.usecase.GetAllUsers()
	// fmt.Printf("%+v", response)
	if (response.Status != "ok") {
		c.JSON(response.StatusCode, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

func (res *userDelivery) GetUserById(c *gin.Context) {
	id := c.Param("id")
	response := res.usecase.GetUserById(id)
	if response.StatusCode == http.StatusNotFound {
		c.JSON(http.StatusOK, response)
		return
	}

	if (response.Status != "ok") {
		c.JSON(response.StatusCode, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

func (res *userDelivery) CreateNewUser(c *gin.Context) {
	request := dto.User{}
	if err := c.ShouldBindJSON(&request); err != nil {
		errorRes := helpers.ResponseError("Bad Request", err, 400)
		c.JSON(errorRes.StatusCode, errorRes)
		return
	}
	response := res.usecase.CreateNewUser(request)

	if (response.Status != "ok") {
		c.JSON(response.StatusCode, response)
		return
	}
	c.JSON(http.StatusOK, response)	
}	

func (res *userDelivery) UpdateUserData(c *gin.Context) {
  	id := c.Param("id")
	request := dto.User{}
	if err := c.ShouldBindJSON(&request); err != nil {
		errorRes := helpers.ResponseError("Bad Request", err, 400)
		c.JSON(errorRes.StatusCode, errorRes)
		return
	}

	response := res.usecase.UpdateUserData(request, id)
	
	if (response.Status != "ok") {
		c.JSON(response.StatusCode, response)
		return
	}
	
	c.JSON(response.StatusCode, response)
}

func (res *userDelivery) DeleteUserById(c *gin.Context) {
	id := c.Param("id")
	response := res.usecase.DeleteUserById(id)
	if (response.Status != "ok") {
		c.JSON(response.StatusCode, response)
		return
	}

	c.JSON(response.StatusCode, response)
}
