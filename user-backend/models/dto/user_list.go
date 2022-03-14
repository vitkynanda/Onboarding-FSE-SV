package dto

type UserList struct {
	Id     string `json:"id"`
	Name   string `json:"name" validate:"required"`
	Role   Role   `json:"role"  validate:"required"`
	Active bool   `json:"active" validate:"required"`
}