package dto

type UpdateAccountInputDTO struct {
	ID      string  `json:"id" validate:"required"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Balance float64 `json:"balance"`
	APIKey  string  `json:"api_key"`
}
