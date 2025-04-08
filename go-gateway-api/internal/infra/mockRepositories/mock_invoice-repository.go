package mockRepositories

import (
	"errors"
	"fmt"

	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/domain/domainEntities"
)

type InMemoryInvoiceRepository struct {
	invoices []*domainEntities.InvoiceDomain
}

func NewInMemoryInvoiceRepository() *InMemoryInvoiceRepository {
	return &InMemoryInvoiceRepository{
		invoices: make([]*domainEntities.InvoiceDomain, 0),
	}
}

func (repo *InMemoryInvoiceRepository) CreateInvoice(invoice *domainEntities.InvoiceDomain) error {
	repo.invoices = append(repo.invoices, invoice)
	return nil
}

func (repo *InMemoryInvoiceRepository) FindByID(id string) (*domainEntities.InvoiceDomain, error) {
	fmt.Println("Finding invoice by ID:", id)
	fmt.Println("Invoices in repository:", repo.invoices)
	fmt.Println("Invoices in repository count:", len(repo.invoices))
	for _, invoice := range repo.invoices {
		if invoice.ID == id {
			return invoice, nil
		}
	}
	return nil, errors.New("invoice not found")
}

func (repo *InMemoryInvoiceRepository) FindByAccountID(accountID string) ([]*domainEntities.InvoiceDomain, error) {
	var result []*domainEntities.InvoiceDomain

	for _, invoice := range repo.invoices {
		if invoice.AccountID == accountID {
			result = append(result, invoice)
		}
	}

	return result, nil
}

func (repo *InMemoryInvoiceRepository) UpdateStatus(invoice *domainEntities.InvoiceDomain) error {
	for i, inv := range repo.invoices {
		if inv.ID == invoice.ID {
			repo.invoices[i].Status = invoice.Status
			return nil
		}
	}
	return errors.New("invoice not found")
}
