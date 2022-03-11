package entity

import (
	"github.com/google/uuid"
)

type User struct {
	Id       		uuid.UUID 	`json:"id" gorm:"primaryKey"`
	Personal_number string 		`json:"personal_number" binding:"required"`
	Name  			string   	`json:"name" binding:"required"`
	Password    	string  	`json:"password" binding:"required"`
	RoleId 			uuid.UUID   `json:"roleId" binding:"required" gorm:"foreignKey"`	
	Email     		string  	`json:"email" binding:"required"`
	Active   		bool    	`json:"active" binding:"required"`
}