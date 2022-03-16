package dto

type UserLogin struct {
	PersonalNumber string `json:"personalNumber" binding:"required"`
	Password       string `json:"password" binding:"required"`
}