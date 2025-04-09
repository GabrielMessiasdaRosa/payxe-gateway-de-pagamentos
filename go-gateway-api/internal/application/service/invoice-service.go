package service

import (
	"fmt"

	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/application/dto"
	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/domain/domainRepositories"
)

type InvoiceService struct {
	InvoiceRepository domainRepositories.InvoiceDomainRepository
	AccountService    *AccountService
}

func NewInvoiceService(invoiceRepository domainRepositories.InvoiceDomainRepository, accountService *AccountService) *InvoiceService {
	return &InvoiceService{
		InvoiceRepository: invoiceRepository,
		AccountService:    accountService,
	}
}

func (i *InvoiceService) CreateInvoice(newInvoice dto.CreateInvoiceInputDTO, apiKey string) (*dto.InvoiceOutputDTO, error) {
	// Verify if account exists
	account, err := i.AccountService.FindByAPIKey(apiKey)
	if err != nil {
		return nil, fmt.Errorf("account not found: %w", err)
	}
	newInvoiceDomain := dto.ToInvoiceDomain(newInvoice, account.ID)
	if newInvoiceDomain == nil {
		return nil, fmt.Errorf("failed to create invoice domain")
	}
	err = i.InvoiceRepository.CreateInvoice(newInvoiceDomain)
	if err != nil {
		return nil, fmt.Errorf("failed to save invoice: %w", err)
	}

	output := dto.FromInvoice(newInvoiceDomain)
	return output, nil
}

func (i *InvoiceService) FindInvoiceByID(id string, apiKey string) (*dto.InvoiceOutputDTO, error) {
	invoice, err := i.InvoiceRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	account, err := i.AccountService.FindByAPIKey(apiKey)
	if err != nil {
		return nil, fmt.Errorf("account not found: %w", err)
	}
	if invoice.AccountID != account.ID {
		return nil, fmt.Errorf("invoice does not belong to the account")
	}

	output := dto.FromInvoice(invoice)
	return output, nil
}

func (i *InvoiceService) ListInvoicesByAccount(accountID string) ([]*dto.InvoiceOutputDTO, error) {
	// Verify if account exists
	_, err := i.AccountService.FindByID(accountID)
	if err != nil {
		return nil, fmt.Errorf("account not found: %w", err)
	}

	// Get invoices
	invoices, err := i.InvoiceRepository.FindByAccountID(accountID)
	if err != nil {
		return nil, err
	}

	// Convert to DTOs
	outputInvoices := make([]*dto.InvoiceOutputDTO, len(invoices))
	for i, invoice := range invoices {
		outputInvoices[i] = dto.FromInvoice(invoice)
	}

	return outputInvoices, nil
}
