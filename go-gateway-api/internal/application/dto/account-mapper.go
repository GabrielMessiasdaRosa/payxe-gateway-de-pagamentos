package dto

import (
	"fmt"

	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/domain/domainEntities"
)

func FromAccount(account *domainEntities.AccountDomain) *AccountOutputDTO {
	fmt.Printf("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX %v\n", account)
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
	fmt.Printf("Account ID: %v\n", account.ID) // should print account.ID as value of account.ID in a fmt.Printf
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
