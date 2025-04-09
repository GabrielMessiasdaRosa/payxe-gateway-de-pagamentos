package mockRepositories

import (
	"testing"

	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/domain/domainEntities"
	"github.com/stretchr/testify/assert"
)

func TestNewInMemoryInvoiceRepository(t *testing.T) {
	repo := NewInMemoryInvoiceRepository()
	assert.NotNil(t, repo)
	assert.Empty(t, repo.invoices)
}

func TestCreateInvoice(t *testing.T) {
	repo := NewInMemoryInvoiceRepository()
	invoice := &domainEntities.InvoiceDomain{
		ID:          "1",
		AccountID:   "acc-1",
		Amount:      100.0,
		Status:      "pending",
		Description: "Test invoice",
	}

	err := repo.CreateInvoice(invoice)
	assert.Nil(t, err)
	assert.Len(t, repo.invoices, 1)
	assert.Equal(t, invoice, repo.invoices[0])
}

func TestFindInvoiceByID(t *testing.T) {
	repo := NewInMemoryInvoiceRepository()
	invoice1 := &domainEntities.InvoiceDomain{
		ID:          "1",
		AccountID:   "acc-1",
		Amount:      100.0,
		Status:      "pending",
		Description: "Test invoice 1",
	}
	invoice2 := &domainEntities.InvoiceDomain{
		ID:          "2",
		AccountID:   "acc-2",
		Amount:      200.0,
		Status:      "pending",
		Description: "Test invoice 2",
	}

	repo.CreateInvoice(invoice1)
	repo.CreateInvoice(invoice2)

	t.Run("should find invoice by ID", func(t *testing.T) {
		found, err := repo.FindByID("1")
		assert.Nil(t, err)
		assert.Equal(t, invoice1, found)

		found, err = repo.FindByID("2")
		assert.Nil(t, err)
		assert.Equal(t, invoice2, found)
	})

	t.Run("should return error when invoice not found", func(t *testing.T) {
		found, err := repo.FindByID("non-existent")
		assert.Error(t, err)
		assert.Equal(t, "invoice not found", err.Error())
		assert.Nil(t, found)
	})
}

func TestFindInvoicesByAccountID(t *testing.T) {
	repo := NewInMemoryInvoiceRepository()
	invoice1 := &domainEntities.InvoiceDomain{
		ID:          "1",
		AccountID:   "acc-1",
		Amount:      100.0,
		Status:      "pending",
		Description: "Test invoice 1",
	}
	invoice2 := &domainEntities.InvoiceDomain{
		ID:          "2",
		AccountID:   "acc-1",
		Amount:      200.0,
		Status:      "paid",
		Description: "Test invoice 2",
	}
	invoice3 := &domainEntities.InvoiceDomain{
		ID:          "3",
		AccountID:   "acc-2",
		Amount:      300.0,
		Status:      "pending",
		Description: "Test invoice 3",
	}

	repo.CreateInvoice(invoice1)
	repo.CreateInvoice(invoice2)
	repo.CreateInvoice(invoice3)

	t.Run("should find invoices by account ID", func(t *testing.T) {
		found, err := repo.FindByAccountID("acc-1")
		assert.Nil(t, err)
		assert.Len(t, found, 2)
		assert.Contains(t, found, invoice1)
		assert.Contains(t, found, invoice2)

		found, err = repo.FindByAccountID("acc-2")
		assert.Nil(t, err)
		assert.Len(t, found, 1)
		assert.Contains(t, found, invoice3)
	})

	t.Run("should return empty slice when no invoices found", func(t *testing.T) {
		found, err := repo.FindByAccountID("non-existent")
		assert.Nil(t, err)
		assert.Empty(t, found)
	})
}

func TestUpdateInvoiceStatus(t *testing.T) {
	repo := NewInMemoryInvoiceRepository()
	invoice := &domainEntities.InvoiceDomain{
		ID:          "1",
		AccountID:   "acc-1",
		Amount:      100.0,
		Status:      "pending",
		Description: "Test invoice",
	}

	repo.CreateInvoice(invoice)

	t.Run("should update invoice status", func(t *testing.T) {
		updatedInvoice := &domainEntities.InvoiceDomain{
			ID:          "1",
			AccountID:   "acc-1",
			Amount:      100.0,
			Status:      "paid",
			Description: "Test invoice",
		}

		err := repo.UpdateStatus(updatedInvoice)
		assert.Nil(t, err)

		found, _ := repo.FindByID("1")
		assert.Equal(t, updatedInvoice.Status, found.Status)
	})

	t.Run("should return error when invoice not found", func(t *testing.T) {
		nonExistentInvoice := &domainEntities.InvoiceDomain{
			ID:     "non-existent",
			Status: "paid",
		}

		err := repo.UpdateStatus(nonExistentInvoice)
		assert.Error(t, err)
		assert.Equal(t, "invoice not found", err.Error())
	})
}
