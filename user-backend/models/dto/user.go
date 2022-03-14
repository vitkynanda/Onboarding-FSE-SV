package dto

type User struct {
	Id              string `json:"id"`
	Name            string `json:"name" validate:"required"`
	Personal_number string `json:"personalNumber" validate:"required"`
	Email           string `json:"email" validate:"required"`
	Role            Role   `json:"role"`
	Active          bool   `json:"active" validate:"required"`
	Password        string `json:"password" validate:"required"`
}