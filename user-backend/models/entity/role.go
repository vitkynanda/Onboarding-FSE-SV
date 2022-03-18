package entity

type Role struct {
	ID     string `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Active bool   `json:"active"`
}