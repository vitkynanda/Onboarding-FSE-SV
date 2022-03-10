package dto

import "github.com/google/uuid"

type UserList struct {
	Id         uuid.UUID `json:"id"`
	Name  	   string    `json:"name" binding:"required"`
	Role  	   Role    	 `json:"role"  binding:"required"`
	Active     bool      `json:"active" binding:"required"`
}