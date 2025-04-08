package service

import (
	"fmt"

	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/application/dto"
	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/domain/domainEntities"
	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/domain/domainRepositories"
)

type AccountService struct {
	accountRepository domainRepositories.AccountDomainRepository
}

func NewAccountService(accountRepository domainRepositories.AccountDomainRepository) *AccountService {
	return &AccountService{
		accountRepository: accountRepository,
	}
}

func (accService *AccountService) CreateAccount(acc *dto.CreateAccountInputDTO) (*dto.AccountOutputDTO, error) {
	newAccount := domainEntities.NewAccount(acc.Name, acc.Email)
	err := accService.accountRepository.CreateAccount(newAccount)
	if err != nil {
		return nil, err
	}
	output := dto.FromAccount(newAccount)
	return output, nil

}

func (accService *AccountService) FindByAPIKey(apiKey string) (*dto.AccountOutputDTO, error) {
	account, err := accService.accountRepository.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, domainEntities.ErrAccountNotFound
	}
	output := dto.FromAccount(account)
	return output, nil
}

func (accService *AccountService) FindByID(id string) (*dto.AccountOutputDTO, error) {
	account, err := accService.accountRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, domainEntities.ErrAccountNotFound
	}
	output := dto.FromAccount(account)
	return output, nil
}

func (accService *AccountService) UpdateBalance(acc *dto.UpdateAccountInputDTO) error {
	fmt.Printf("WWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWW %v\n", acc)
	account, err := accService.accountRepository.FindByID(acc.ID)
	fmt.Printf("Account ID: %v\n", account)
	if err != nil {
		return err
	}
	if account == nil {
		return domainEntities.ErrAccountNotFound
	}
	return accService.accountRepository.UpdateBalance(account)
}
