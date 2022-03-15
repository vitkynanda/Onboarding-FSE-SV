package dto

type User struct {
	Id              string `json:"id"`
	Name            string `json:"name" binding:"required" gorm:"type=varchar(255)"`
	Personal_number string `json:"personalNumber" binding:"required" gorm:"type=varchar(255)"`
	Email           string `json:"email" binding:"required" gorm:"type=varchar(255)"`
	Role            Role   `json:"role"`
	Active          bool   `json:"active"`
	Password        string `json:"password" binding:"required" gorm:"type=varchar(255)"`
}