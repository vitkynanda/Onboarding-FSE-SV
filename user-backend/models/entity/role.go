package entity

type Role struct {
	Id     string `json:"id" binding:"required"`
	Title  string `json:"title" binding:"required"`
	Active string `json:"active" binding:"required"`
}