package dto

import (
	"fmt"
	"time"

	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/domain/domainEntities"
)

func FromInvoice(invoice *domainEntities.InvoiceDomain) *InvoiceOutputDTO {
	if invoice == nil {
		panic("Invoice is nil")
	}

	output := InvoiceOutputDTO{
		ID:          invoice.ID,
		Amount:      invoice.Amount,
		Description: invoice.Description,
		Status:      string(invoice.Status),
		AccountID:   invoice.AccountID,
		CreatedAt:   invoice.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   invoice.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	return &output
}

func ToInvoiceDomain(dto interface{}) *domainEntities.InvoiceDomain {
	if dto == nil {
		return nil
	}

	var invoice *domainEntities.InvoiceDomain

	switch v := dto.(type) {
	case CreateInvoiceInputDTO:
		invoice = &domainEntities.InvoiceDomain{
			Amount:      v.Amount,
			Description: v.Description,
			AccountID:   v.AccountID,
		}
	case UpdateInvoiceInputDTO:
		invoice = &domainEntities.InvoiceDomain{
			ID:          v.ID,
			Amount:      v.Amount,
			Description: v.Description,
			Status:      domainEntities.Status(v.Status),
			AccountID:   v.AccountID,
		}
	case InvoiceOutputDTO:
		createdAt, _ := time.Parse("2006-01-02 15:04:05", v.CreatedAt)
		updatedAt, _ := time.Parse("2006-01-02 15:04:05", v.UpdatedAt)
		invoice = &domainEntities.InvoiceDomain{
			ID:          v.ID,
			Amount:      v.Amount,
			Description: v.Description,
			Status:      domainEntities.Status(v.Status),
			AccountID:   v.AccountID,
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
		}
	default:
		fmt.Printf("Unsupported type: %T\n", v)
		return nil
	}

	return invoice
}
