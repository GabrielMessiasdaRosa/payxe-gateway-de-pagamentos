package mockRepositories

import (
	"errors"

	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/domain/domainEntities"
)

type InMemoryAccountRepository struct {
	accounts []*domainEntities.AccountDomain
}

func NewInMemoryAccountRepository() *InMemoryAccountRepository {
	return &InMemoryAccountRepository{
		accounts: make([]*domainEntities.AccountDomain, 0),
	}
}

func (repo *InMemoryAccountRepository) CreateAccount(account *domainEntities.AccountDomain) error {
	repo.accounts = append(repo.accounts, account)
	return nil
}

func (repo *InMemoryAccountRepository) FindByID(id string) (*domainEntities.AccountDomain, error) {
	for _, acc := range repo.accounts {
		if acc.ID == id {
			return acc, nil
		}
	}
	return nil, errors.New("account not found")
}

func (repo *InMemoryAccountRepository) FindByAPIKey(apiKey string) (*domainEntities.AccountDomain, error) {
	for _, acc := range repo.accounts {
		if acc.APIKey == apiKey {
			return acc, nil
		}
	}
	return nil, errors.New("account not found")
}

func (repo *InMemoryAccountRepository) UpdateBalance(account *domainEntities.AccountDomain) error {
	for i, acc := range repo.accounts {
		if acc.ID == account.ID {
			repo.accounts[i] = account
			return nil
		}
	}
	return errors.New("account not found")
}
