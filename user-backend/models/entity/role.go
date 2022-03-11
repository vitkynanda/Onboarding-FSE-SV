package entity

type Role struct {
	ID     string `json:"id" gorm:"primaryKey, type:varchar(50)"`
	Title  string `json:"title" `
	Active bool   `json:"active" binding:"required"`
}