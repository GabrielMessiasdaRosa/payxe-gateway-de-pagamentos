package dto

import (
	"fmt"

	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/domain/domainEntities"
)

func FromAccount(account *domainEntities.AccountDomain) *AccountOutputDTO {
	if account == nil {
		panic("Account is nil")
	}

	output := AccountOutputDTO{
		ID:      account.ID,
		Name:    account.Name,
		Email:   account.Email,
		Balance: account.Balance,
		APIKey:  account.APIKey,
	}
	return &output
}

func ToAccountDomain(dto interface{}) *domainEntities.AccountDomain {
	account := &domainEntities.AccountDomain{}
	if dto == nil {
		fmt.Println("DTO is nil")
		return nil
	}
	switch v := dto.(type) {
	case CreateAccountInputDTO:
		account.Name = v.Name
		account.Email = v.Email
	case UpdateAccountInputDTO:
		account.ID = v.ID
		account.Balance = v.Balance
		account.APIKey = v.APIKey
		account.Name = v.Name
		account.Email = v.Email
	case AccountOutputDTO:
		account.ID = v.ID
		account.Name = v.Name
		account.Email = v.Email
		account.Balance = v.Balance
		account.APIKey = v.APIKey
	default:
		fmt.Printf("Unsupported type: %T\n", v)
		return nil
	}

	return account
}
