package dto

import (
	"testing"
	"time"

	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/domain/domainEntities"
	"github.com/stretchr/testify/assert"
)

func TestFromInvoice(t *testing.T) {
	t.Run("should convert invoice domain to dto", func(t *testing.T) {
		now := time.Now()
		invoice := &domainEntities.InvoiceDomain{
			ID:             "123",
			Amount:         100.0,
			Description:    "Test invoice",
			Status:         domainEntities.StatusApproved,
			AccountID:      "acc123",
			CardLastDigits: "1234",
			CreatedAt:      now,
			UpdatedAt:      now,
		}

		output := FromInvoice(invoice)

		assert.Equal(t, invoice.ID, output.ID)
		assert.Equal(t, invoice.Amount, output.Amount)
		assert.Equal(t, invoice.Description, output.Description)
		assert.Equal(t, string(invoice.Status), output.Status)
		assert.Equal(t, invoice.AccountID, output.AccountID)
		assert.Equal(t, invoice.CreatedAt.Format("2006-01-02 15:04:05"), output.CreatedAt)
		assert.Equal(t, invoice.UpdatedAt.Format("2006-01-02 15:04:05"), output.UpdatedAt)
	})

	t.Run("should panic when invoice is nil", func(t *testing.T) {
		assert.Panics(t, func() {
			FromInvoice(nil)
		})
	})
}

func TestToInvoiceDomain(t *testing.T) {
	t.Run("should convert CreateInvoiceInputDTO to domain", func(t *testing.T) {
		input := CreateInvoiceInputDTO{
			Amount:          150.0,
			Description:     "New invoice",
			CardNumber:      "4111111111111111",
			CVV:             "123",
			CardHolderName:  "John Doe",
			ExpirationMonth: 12,
			ExpirationYear:  2030,
		}
		accountID := "acc456"

		result := ToInvoiceDomain(input, accountID)
		assert.NotNil(t, result)
		assert.NotEmpty(t, result.ID)
		assert.Equal(t, input.Amount, result.Amount)
		assert.Equal(t, input.Description, result.Description)
		assert.Equal(t, accountID, result.AccountID)
		assert.Equal(t, domainEntities.StatusPending, result.Status)
		assert.Equal(t, "1111", result.CardLastDigits)
	})

	t.Run("should convert UpdateInvoiceInputDTO to domain", func(t *testing.T) {
		input := UpdateInvoiceInputDTO{
			ID:          "789",
			Amount:      200.0,
			Description: "Updated invoice",
			Status:      string(domainEntities.StatusApproved),
		}
		accountID := "acc789"

		result := ToInvoiceDomain(input, accountID)

		assert.NotNil(t, result)
		assert.Equal(t, input.ID, result.ID)
		assert.Equal(t, input.Amount, result.Amount)
		assert.Equal(t, input.Description, result.Description)
		assert.Equal(t, domainEntities.Status(input.Status), result.Status)
		assert.Equal(t, accountID, result.AccountID)
	})

	t.Run("should convert InvoiceOutputDTO to domain", func(t *testing.T) {
		now := time.Now().Format("2006-01-02 15:04:05")
		input := InvoiceOutputDTO{
			ID:          "456",
			Amount:      300.0,
			Description: "Existing invoice",
			Status:      string(domainEntities.StatusApproved),
			AccountID:   "acc123",
			CreatedAt:   now,
			UpdatedAt:   now,
		}
		accountID := "acc123"

		result := ToInvoiceDomain(input, accountID)

		assert.NotNil(t, result)
		assert.Equal(t, input.ID, result.ID)
		assert.Equal(t, input.Amount, result.Amount)
		assert.Equal(t, input.Description, result.Description)
		assert.Equal(t, domainEntities.Status(input.Status), result.Status)
		assert.Equal(t, input.AccountID, result.AccountID)

		parsedTime, _ := time.Parse("2006-01-02 15:04:05", now)
		assert.Equal(t, parsedTime, result.CreatedAt)
		assert.Equal(t, parsedTime, result.UpdatedAt)
	})

	t.Run("should return nil for invalid input type", func(t *testing.T) {
		result := ToInvoiceDomain("invalid input", "acc123")
		assert.Nil(t, result)
	})

	t.Run("should return nil for invalid credit card", func(t *testing.T) {
		input := CreateInvoiceInputDTO{
			Amount:          150.0,
			Description:     "New invoice",
			CardNumber:      "invalid",
			CVV:             "123",
			CardHolderName:  "John Doe",
			ExpirationMonth: 12,
			ExpirationYear:  2030,
		}

		result := ToInvoiceDomain(input, "acc123")
		assert.Nil(t, result)
	})
}
