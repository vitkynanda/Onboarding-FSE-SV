package dto

import "github.com/google/uuid"

type Role struct {
	Id    uuid.UUID `json:"id" `
	Title string    `json:"title" `
}