package domainRepositories

import "github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/domain/domainEntities"

type InvoiceDomainRepository interface {
	CreateInvoice(invoice *domainEntities.InvoiceDomain) error
	FindByID(id string) (*domainEntities.InvoiceDomain, error)
	UpdateStatus(invoice *domainEntities.InvoiceDomain) error
}
