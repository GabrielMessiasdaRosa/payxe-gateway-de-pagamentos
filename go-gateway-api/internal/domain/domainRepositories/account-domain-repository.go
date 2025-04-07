package domainRepositories

import "github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/domain/domainEntities"


type AccountDomainRepository interface {
	Save(account *domainEntities.AccountDomain) error
	FindByID(id string) (*domainEntities.AccountDomain, error)
	FindByAPIKey(apiKey string) (*domainEntities.AccountDomain, error)
	UpdateBalance(account *domainEntities.AccountDomain) error
	DeleteByID(id string) error
	ListAll() ([]*domainEntities.AccountDomain, error) 
}