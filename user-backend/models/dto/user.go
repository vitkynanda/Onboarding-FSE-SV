package dto

import "github.com/google/uuid"


type User struct {
	Id         		uuid.UUID `json:"id"`
	Name  	   		string    `json:"name" binding:"required"`
	Personal_number string    `json:"personalNumber"  binding:"required"`
	Email     		string    `json:"email" binding:"required"`
	Role  			Role      `json:"role"`
	Active     		bool      `json:"active"`
	Password        string 	  `json:"password" binding:"required"`
}