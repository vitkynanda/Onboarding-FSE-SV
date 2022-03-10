package dto

type Role struct {
	Id    string `json:"id" binding:"required"`
	Title string `json:"title" binding:"required"`
}