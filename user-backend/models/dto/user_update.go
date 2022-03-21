package dto

type RoleUpdate struct {
	Id string `json:"id"`
}
type UserUpdate struct {
	Id              string     `json:"id"`
	Name            string     `json:"name"`
	Personal_number string     `json:"personalNumber"`
	Email           string     `json:"email"`
	Role            RoleUpdate `json:"role"`
	Active          bool       `json:"active"`
	Password        string     `json:"password"`
}