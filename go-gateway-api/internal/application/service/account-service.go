package service

import (
	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/application/dto"
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

func (s *AccountService) Save(acc *dto.CreateAccountInputDTO) (*dto.AccountOutputDTO, error) {
	account := dto.ToAccount(*acc)
	err := s.accountRepository.Save(account)
	if err != nil {
		return nil, err
	}
	output := dto.FromAccount(account)
	return &output, nil
	
}



