package dto

type UpdateApplicationRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
	Type        string `json:"type"`
}
