package dto

import "github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/domain"

type CreateAccountInputDTO struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type UpdateAccountInputDTO struct {
	Name  *string `json:"name" validate:"omitempty"`
	Email *string `json:"email" validate:"omitempty,email"`
}

type AccountOutputDTO struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Balance float64 `json:"balance"`
	APIKey  string  `json:"api_key"`
}


func ToAccount(input CreateAccountInputDTO) *domain.Account {
	return domain.NewAccount(input.Name, input.Email)
}

func FromAccount(account *domain.Account) AccountOutputDTO {
	return AccountOutputDTO{
		ID:      account.ID,
		Name:    account.Name,
		Email:   account.Email,
		Balance: account.Balance,
		APIKey:  account.APIKey,
	}
}