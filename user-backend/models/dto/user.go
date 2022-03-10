package dto

import "github.com/google/uuid"


type User struct {
	Id         		uuid.UUID `json:"id"`
	Name  	   		string    `json:"name" binding:"required"`
	Personal_number string    `json:"personal_number"  binding:"required"`
	Email     		string    `json:"gender" binding:"required"`
	Role  			Role      `json:"birthdate" binding:"required"`
	Active     		bool      `json:"active" binding:"required"`
}