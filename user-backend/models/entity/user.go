package entity

type User struct {
	ID              string `gorm:"primaryKey"`
	Personal_number string
	Name            string
	Password        string
	Email           string
	RoleID          string
	Active          bool
}