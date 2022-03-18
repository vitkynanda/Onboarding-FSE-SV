package dto

type UserList struct {
	Id     string      `json:"id"`
	Name   string      `json:"name" validate:"required"`
	Role   interface{} `json:"role"  validate:"required"`
	Active bool        `json:"active" validate:"required"`
}