package user_delivery

import (
	"fmt"
	"go-api/helpers"
	"go-api/models/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (res *userDelivery) GetAllUsers(c *gin.Context) {
	response := res.usecase.GetAllUsers()
	if (response.StatusCode != 200) {
		c.JSON(http.StatusInternalServerError, gin.H{"message" : "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (res *userDelivery) GetUserById(c *gin.Context) {
	id := c.Param("id")
	response := res.usecase.GetUserById(id)
	if (response.StatusCode != 200) {
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (res *userDelivery) CreateNewUser(c *gin.Context) {
	request := dto.User{}
	if err := c.ShouldBindJSON(&request); err != nil {
		errorMessages :=  []string{}
	
		for _, e :=  range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on Field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages,  errorMessage)
		} 
		
		if len(errorMessages) > 0 {
			errorRes := helpers.ResponseError("Invalid Input", 400)

			c.JSON(http.StatusBadRequest, errorRes)
			return
		}

	}

	response := res.usecase.CreateNewUser(request)

	if (response.StatusCode != 200) {
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	c.JSON(http.StatusOK, response)
	
}	

func (res *userDelivery) UpdateUserData(c *gin.Context) {
  	id := c.Param("id")
	request := dto.User{}

	if err := c.ShouldBindJSON(&request); err != nil {
		errorMessages :=  []string{}
	
		fmt.Printf("%+v", request)
		for _, e :=  range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on Field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages,  errorMessage)
			
		} 
		
		if len(errorMessages) > 0 {
			errorRes := helpers.ResponseError("Invalid Input", 400)
			c.JSON(http.StatusBadRequest, errorRes)
			return
		}

	}

	response := res.usecase.UpdateUserData(request, id)

	if (response.StatusCode != 200) {
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	
	c.JSON(http.StatusOK, response)
}

func (res *userDelivery) DeleteUserById(c *gin.Context) {
	id := c.Param("id")
	response := res.usecase.DeleteUserById(id)
	if (response.StatusCode != 200) {
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	c.JSON(http.StatusOK, response)
}
