package dto

type UserLogin struct {
	PersonalNumber string `json:"personal_number"`
	Password       string `json:"password"`
}