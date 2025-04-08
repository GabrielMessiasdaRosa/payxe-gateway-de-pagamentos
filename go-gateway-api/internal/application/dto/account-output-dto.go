package dto

type AccountOutputDTO struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Balance float64 `json:"balance"`
	APIKey  string  `json:"api_key"`
}
