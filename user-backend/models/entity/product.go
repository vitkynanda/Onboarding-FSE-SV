package entity

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID          string
	Name        string
	Description string
	Status      string
	MakerID     string
	SignerID    string
	CheckerID   string
	UpdatedAt 	time.Time
}
