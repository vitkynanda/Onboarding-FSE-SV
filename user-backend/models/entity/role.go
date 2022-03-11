package entity

type Role struct {
	ID     string `gorm:"primaryKey"`
	Title  string
	Active bool
}