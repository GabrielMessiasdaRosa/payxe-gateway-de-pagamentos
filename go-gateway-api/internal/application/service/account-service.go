package service

import (
	"fmt"

	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/application/dto"
	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/infra/repositories"
)

type AccountService struct {
	repository *repositories.AccountRepository
}

func NewAccountService(repository *repositories.AccountRepository) *AccountService {
	return &AccountService{
		repository: repository,
	}
}

func (accService *AccountService) Save(input dto.CreateAccountInputDTO) (*dto.AccountOutputDTO, error) {
	account := dto.ToAccount(input)

	existingAccount, err := accService.repository.FindByAPIKey(account.APIKey)
	if err != nil {
		return nil, err
	}

	if existingAccount != nil {
		return nil, fmt.Errorf("account with API key %s already exists", account.APIKey)
	}

	err = accService.repository.Save(account)
	if err != nil {
		return nil, err
	}
	// Convert the account to the output DTO
	dtoAccount := dto.FromAccount(account)

	return &dtoAccount, nil
}

func (accService *AccountService) UpdateBalance(apiKey string, amount float64) (*dto.AccountOutputDTO, error) {
	account, err := accService.repository.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, fmt.Errorf("account with API key %s not found", apiKey)
	}

	account.AddBalance(amount)
	err = accService.repository.UpdateBalance(account)
	if err != nil {
		return nil, err
	}

	dtoAccount := dto.FromAccount(account)
	return &dtoAccount, nil
}

func (accServce *AccountService) FindByAPIKey(apiKey string) (*dto.AccountOutputDTO, error) {
	account, err := accServce.repository.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, fmt.Errorf("account with API key %s not found", apiKey)
	}

	dtoAccount := dto.FromAccount(account)
	return &dtoAccount, nil
}

func (accService *AccountService) FindByID(id string) (*dto.AccountOutputDTO, error) {
	account, err := accService.repository.FindByID(id)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, fmt.Errorf("account with ID %s not found", id)
	}
	dtoAccount := dto.FromAccount(account)
	return &dtoAccount, nil
}
	
