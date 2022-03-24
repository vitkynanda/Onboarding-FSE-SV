package dto

type Product struct {
	ID          string `json:"id"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
	MakerID     string `json:"marker_id"`
	SignerID    string `json:"signer_id"`
	CheckerID   string `json:"checker_id"`
}