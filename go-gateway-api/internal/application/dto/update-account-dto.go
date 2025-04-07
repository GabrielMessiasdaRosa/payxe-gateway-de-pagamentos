package dto

type UpdateAccountInputDTO struct {
	Name  *string `json:"name" validate:"omitempty"`
	Email *string `json:"email" validate:"omitempty,email"`
	APIKey *string `json:"api_key" validate:"omitempty"`
}

type UpdateAccountOutputDTO struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Balance float64 `json:"balance"`
	APIKey  string  `json:"api_key"`
}