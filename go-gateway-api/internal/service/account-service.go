package domain

import (
	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/domain"
)

type AccountService struct {
	repository *domain.AccountRepository
}

func NewAccountService(repository *domain.AccountRepository) *AccountService {
	return &AccountService{
		repository: repository,
	}
}