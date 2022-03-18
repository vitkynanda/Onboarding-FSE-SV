package dto

type Role struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Active bool   `json:"active"`
}