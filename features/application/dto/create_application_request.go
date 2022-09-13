package dto

type CreateApplicationRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Type        string `json:"type" binding:"omitempty"`
	Enabled     bool   `json:"enabled" binding:"omitempty"`
}
