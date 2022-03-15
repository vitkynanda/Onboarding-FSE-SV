package entity

import "time"

type Product struct {
	ID          string
	Name        string
	Description string
	Status      string
	MakerID    string
	SignerID    string
	CheckerID   string
	CreatedAt   time.Time
	UpdatedAt 	time.Time
	DeletedAt 	time.Time
}
