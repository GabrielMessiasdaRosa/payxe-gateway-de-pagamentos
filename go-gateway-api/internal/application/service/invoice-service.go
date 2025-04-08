// filepath: /home/gmrosa/Desktop/Estudo/imersoes/payxe-gateway-de-pagamentos/go-gateway-api/internal/application/service/invoice-service.go
package service

import (
	"fmt"

	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/application/dto"
	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/domain/domainEntities"
	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/domain/domainRepositories"
	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/domain/valueObjects"
)

type InvoiceService struct {
	invoiceRepository domainRepositories.InvoiceDomainRepository
	accountRepository domainRepositories.AccountDomainRepository
}

func NewInvoiceService(invoiceRepository domainRepositories.InvoiceDomainRepository, accountRepository domainRepositories.AccountDomainRepository) *InvoiceService {
	return &InvoiceService{
		invoiceRepository: invoiceRepository,
		accountRepository: accountRepository,
	}
}

func (i *InvoiceService) CreateInvoice(newInvoice *dto.CreateInvoiceInputDTO) (*dto.InvoiceOutputDTO, error) {
	// Verify if account exists
	account, err := i.accountRepository.FindByAPIKey(newInvoice.APIKey)
	if err != nil {
		return nil, fmt.Errorf("account not found: %w", err)
	}
	card, err := valueObjects.NewCreditCard(
		newInvoice.CardNumber,
		newInvoice.CVV,
		newInvoice.CardHolderName,
		newInvoice.ExpirationMonth,
		newInvoice.ExpirationYear,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create credit card: %w", err)
	}
	newInvoiceDomain, err := domainEntities.NewInvoice(
		account.ID,
		newInvoice.Amount,
		newInvoice.Description,
		newInvoice.PaymentType,
		card,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create invoice: %w", err)
	}
	err = i.invoiceRepository.CreateInvoice(newInvoiceDomain)
	if err != nil {
		return nil, fmt.Errorf("failed to save invoice: %w", err)
	}

	output := dto.FromInvoice(newInvoiceDomain)
	return output, nil
}

func (i *InvoiceService) FindInvoiceByID(id string) (*dto.InvoiceOutputDTO, error) {
	invoice, err := i.invoiceRepository.FindByID(id)
	fmt.Println("Invoice:", invoice)
	if err != nil {
		return nil, err
	}

	output := dto.FromInvoice(invoice)
	return output, nil
}

func (i *InvoiceService) ListInvoicesByAccount(accountID string) ([]*dto.InvoiceOutputDTO, error) {
	// Verify if account exists
	_, err := i.accountRepository.FindByID(accountID)
	if err != nil {
		return nil, fmt.Errorf("account not found: %w", err)
	}

	// Get invoices
	invoices, err := i.invoiceRepository.FindByAccountID(accountID)
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
