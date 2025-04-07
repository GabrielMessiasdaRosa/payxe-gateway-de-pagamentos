package dto

import "github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/domain/domainEntities"


type AccountOutputDTO struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Balance float64 `json:"balance"`
	APIKey  string  `json:"api_key"`
}

func ToAccount(input CreateAccountInputDTO) *domainEntities.AccountDomain {
	return domainEntities.NewAccount(input.Name, input.Email)
}

func FromAccount(account *domainEntities.AccountDomain) AccountOutputDTO {
	return AccountOutputDTO{
		ID:      account.ID,
		Name:    account.Name,
		Email:   account.Email,
		Balance: account.Balance,
		APIKey:  account.APIKey,
	}
}