package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID              string `gorm:"primaryKey"`
	Personal_number string
	Name            string
	Password        string
	Email           string
	RoleID          string
	Active          bool
}