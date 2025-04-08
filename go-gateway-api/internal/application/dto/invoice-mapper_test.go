package dto

import (
	"testing"
	"time"

	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/domain/domainEntities"
	"github.com/stretchr/testify/assert"
)

func TestFromInvoice(t *testing.T) {
	t.Run("should convert from invoice domain to invoice output dto", func(t *testing.T) {
		createdAt := time.Now()
		updatedAt := time.Now()

		invoice := &domainEntities.InvoiceDomain{
			ID:          "123",
			Amount:      100.0,
			Description: "Test invoice",
			Status:      "pending",
			AccountID:   "acc123",
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
		}

		output := FromInvoice(invoice)

		assert.Equal(t, invoice.ID, output.ID)
		assert.Equal(t, invoice.Amount, output.Amount)
		assert.Equal(t, invoice.Description, output.Description)
		assert.Equal(t, string(invoice.Status), output.Status)
		assert.Equal(t, invoice.AccountID, output.AccountID)
		assert.Equal(t, createdAt.Format("2006-01-02 15:04:05"), output.CreatedAt)
		assert.Equal(t, updatedAt.Format("2006-01-02 15:04:05"), output.UpdatedAt)
	})

	t.Run("should panic when invoice is nil", func(t *testing.T) {
		assert.Panics(t, func() {
			FromInvoice(nil)
		})
	})
}

func TestToInvoiceDomain(t *testing.T) {
	t.Run("should convert from CreateInvoiceInputDTO to invoice domain", func(t *testing.T) {
		input := CreateInvoiceInputDTO{
			Amount:      100.0,
			Description: "Test invoice",
			AccountID:   "acc123",
		}

		invoice := ToInvoiceDomain(input)

		assert.Equal(t, input.Amount, invoice.Amount)
		assert.Equal(t, input.Description, invoice.Description)
		assert.Equal(t, input.AccountID, invoice.AccountID)
	})

	t.Run("should convert from UpdateInvoiceInputDTO to invoice domain", func(t *testing.T) {
		input := UpdateInvoiceInputDTO{
			ID:          "123",
			Amount:      100.0,
			Description: "Updated invoice",
			Status:      "paid",
			AccountID:   "acc123",
		}

		invoice := ToInvoiceDomain(input)

		assert.Equal(t, input.ID, invoice.ID)
		assert.Equal(t, input.Amount, invoice.Amount)
		assert.Equal(t, input.Description, invoice.Description)
		assert.Equal(t, string(invoice.Status), input.Status)
		assert.Equal(t, input.AccountID, invoice.AccountID)
	})

	t.Run("should convert from InvoiceOutputDTO to invoice domain", func(t *testing.T) {
		input := InvoiceOutputDTO{
			ID:          "123",
			Amount:      100.0,
			Description: "Test invoice",
			Status:      "pending",
			AccountID:   "acc123",
			CreatedAt:   "2023-01-01 12:00:00",
			UpdatedAt:   "2023-01-01 12:30:00",
		}

		invoice := ToInvoiceDomain(input)

		assert.Equal(t, input.ID, invoice.ID)
		assert.Equal(t, input.Amount, invoice.Amount)
		assert.Equal(t, input.Description, invoice.Description)
		assert.Equal(t, input.Status, string(invoice.Status))
		assert.Equal(t, input.AccountID, invoice.AccountID)

		expectedCreatedAt, _ := time.Parse("2006-01-02 15:04:05", input.CreatedAt)
		expectedUpdatedAt, _ := time.Parse("2006-01-02 15:04:05", input.UpdatedAt)

		assert.Equal(t, expectedCreatedAt, invoice.CreatedAt)
		assert.Equal(t, expectedUpdatedAt, invoice.UpdatedAt)
	})

	t.Run("should return nil for unsupported type", func(t *testing.T) {
		invoice := ToInvoiceDomain("unsupported")
		assert.Nil(t, invoice)
	})

	t.Run("should return nil when dto is nil", func(t *testing.T) {
		invoice := ToInvoiceDomain(nil)
		assert.Nil(t, invoice)
	})
}
