package mockRepositories

import (
	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/domain/domainEntities"
	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/domain/domainRepositories"
	"github.com/stretchr/testify/mock"
)

type MockAccountRepository struct {
	accountRepository domainRepositories.AccountDomainRepository
	mock.Mock
}

// DeleteByID implements domainRepositories.AccountDomainRepository.
func (m *MockAccountRepository) DeleteByID(id string) error {
	panic("unimplemented")
}

// ListAll implements domainRepositories.AccountDomainRepository.
func (m *MockAccountRepository) ListAll() ([]*domainEntities.AccountDomain, error) {
	panic("unimplemented")
}

func (m *MockAccountRepository) FindByAPIKey(apiKey string) (*domainEntities.AccountDomain, error) {

	args := m.Called(apiKey)
	return args.Get(0).(*domainEntities.AccountDomain), args.Error(1)
}

func (m *MockAccountRepository) Save(account *domainEntities.AccountDomain) error {

	args := m.Called(account)
	return args.Error(0)
}

func (m *MockAccountRepository) UpdateBalance(account *domainEntities.AccountDomain) error {

	args := m.Called(account)
	return args.Error(0)
}

func (m *MockAccountRepository) FindByID(id string) (*domainEntities.AccountDomain, error) {

	args := m.Called(id)
	return args.Get(0).(*domainEntities.AccountDomain), args.Error(1)
}
