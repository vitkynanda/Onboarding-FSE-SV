package entity

import "time"

type Product struct {
	ID          string
	Name        string
	Description string
	Status      string
	MarkerID    string
	SignerID    string
	CheckerID   string
	CreatedAt   time.Time
	UpdatedAt 	time.Time
	DeletedAt 	time.Time
}
