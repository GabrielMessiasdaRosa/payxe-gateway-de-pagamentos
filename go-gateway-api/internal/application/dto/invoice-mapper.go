package dto

import (
	"fmt"
	"time"

	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/domain/domainEntities"
	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/domain/valueObjects"
)

func FromInvoice(invoice *domainEntities.InvoiceDomain) *InvoiceOutputDTO {
	if invoice == nil {
		panic("Invoice is nil")
	}

	output := &InvoiceOutputDTO{
		ID:          invoice.ID,
		Amount:      invoice.Amount,
		Description: invoice.Description,
		Status:      string(invoice.Status),
		AccountID:   invoice.AccountID,
		CreatedAt:   invoice.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   invoice.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	return output
}

func ToInvoiceDomain(dto any, accountID string) *domainEntities.InvoiceDomain {
	invoice := &domainEntities.InvoiceDomain{}
	fmt.Println("DTO:", dto)
	fmt.Println("invoice:", invoice)
	switch v := dto.(type) {
	case CreateInvoiceInputDTO:
		card, err := valueObjects.NewCreditCard(
			v.CardNumber,
			v.CVV,
			v.CardHolderName,
			v.ExpirationMonth,
			v.ExpirationYear,
		)

		if err != nil {
			fmt.Println("Error creating credit card:", err)
			return nil
		}

		invoice, err = domainEntities.NewInvoice(
			accountID,
			v.Amount,
			v.Description,
			v.PaymentType,
			card,
		)
		fmt.Println("Invoice:", invoice)
		if err != nil {
			fmt.Println("Error creating invoice:", err)
			return nil
		}
		invoice.CardLastDigits = card.GetLastDigits()
		return invoice
	case UpdateInvoiceInputDTO:
		invoice.ID = v.ID
		invoice.Amount = v.Amount
		invoice.Description = v.Description
		invoice.Status = domainEntities.Status(v.Status)
		invoice.AccountID = accountID
		return invoice
	case InvoiceOutputDTO:
		invoice.ID = v.ID
		invoice.Amount = v.Amount
		invoice.Description = v.Description
		invoice.Status = domainEntities.Status(v.Status)
		invoice.AccountID = v.AccountID
		invoice.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", v.CreatedAt)
		invoice.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", v.UpdatedAt)
		return invoice
	default:
		fmt.Printf("Unsupported type: %T\n", v)
		return nil
	}

}
