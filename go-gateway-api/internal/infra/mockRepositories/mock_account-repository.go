package mockRepositories

import (
	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/domain/domainEntities"
	"github.com/stretchr/testify/mock"
)

type MockAccountRepository struct {
	mock.Mock
}

func (m *MockAccountRepository) CreateAccount(account *domainEntities.AccountDomain) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *MockAccountRepository) FindByAPIKey(apiKey string) (*domainEntities.AccountDomain, error) {
	acc := m.Called(apiKey).Get(0)

	// Retorna o valor como *domainEntities.AccountDomain
	account := acc

	if account == nil {
		return nil, nil
	}

	accountDomain, ok := account.(*domainEntities.AccountDomain)
	if !ok {
		panic("interface conversion error: expected *domainEntities.AccountDomain")
	}

	return accountDomain, nil
}

func (m *MockAccountRepository) FindByID(id string) (*domainEntities.AccountDomain, error) {
	ret := m.Called(id)
	// Retorna o valor como *domainEntities.AccountDomain
	account, ok := ret.Get(0).(*domainEntities.AccountDomain)
	if !ok && ret.Get(0) != nil {
		panic("interface conversion error: expected *domainEntities.AccountDomain")
	}
	return account, ret.Error(1)
}

func (m *MockAccountRepository) UpdateBalance(account *domainEntities.AccountDomain) error {
	args := m.Called(account)
	return args.Error(0)
}
