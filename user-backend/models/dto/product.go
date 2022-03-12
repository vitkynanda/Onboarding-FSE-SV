package dto

type Product struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
	MarkerID    string `json:"marker_id"`
	SignerID    string `json:"signer_id"`
	CheckerID   string `json:"checker_id"`
}