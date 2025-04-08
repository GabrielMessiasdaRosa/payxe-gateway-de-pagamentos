package domainEntities

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewInvoice(t *testing.T) {
	accountID := uuid.New()
	amount := float64(100.50)
	description := "Test invoice"
	cardLastDigits := "1234"

	invoice := NewInvoice(accountID, amount, description, cardLastDigits)

	assert.NotNil(t, invoice)
	assert.Equal(t, accountID, invoice.AccountID)
	assert.Equal(t, amount, invoice.Amount)
	assert.Equal(t, "pending", invoice.Status)
	assert.Equal(t, description, invoice.Description)
	assert.Equal(t, cardLastDigits, invoice.CardLastDigits)
	assert.NotEqual(t, uuid.Nil, invoice.ID)
	assert.NotNil(t, invoice.CreatedAt)
	assert.NotNil(t, invoice.UpdatedAt)
}

// @audit-issue
func TestInvoiceSetStatus(t *testing.T) {
	accountID := uuid.New()
	invoice := NewInvoice(accountID, 100.0, "Test invoice", "1234")

	oldUpdatedAt := invoice.UpdatedAt
	time.Sleep(1 * time.Millisecond)

	invoice.SetStatus(StatusPaid)

	assert.Equal(t, StatusPaid, invoice.Status)
	assert.True(t, invoice.UpdatedAt.After(oldUpdatedAt))

	// Test changing to failed status
	oldUpdatedAt = invoice.UpdatedAt
	time.Sleep(1 * time.Millisecond)

	invoice.SetStatus(StatusFailed)

	assert.Equal(t, StatusFailed, invoice.Status)
	assert.True(t, invoice.UpdatedAt.After(oldUpdatedAt))

	// Test changing back to pending
	oldUpdatedAt = invoice.UpdatedAt
	time.Sleep(1 * time.Millisecond)

	invoice.SetStatus(StatusPending)

	assert.Equal(t, StatusPending, invoice.Status)
	assert.True(t, invoice.UpdatedAt.After(oldUpdatedAt))
}

func TestInvoiceIsValid(t *testing.T) {
	tests := []struct {
		name        string
		invoice     *InvoiceDomain
		expectError bool
	}{
		{
			name: "Valid invoice",
			invoice: &InvoiceDomain{
				ID:             uuid.New(),
				AccountID:      uuid.New(),
				Amount:         100.0,
				Status:         "pending",
				Description:    "Test invoice",
				CardLastDigits: "1234",
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			expectError: false,
		},
		{
			name: "Invalid - Empty AccountID",
			invoice: &InvoiceDomain{
				ID:             uuid.New(),
				AccountID:      uuid.Nil,
				Amount:         100.0,
				Status:         "pending",
				Description:    "Test invoice",
				CardLastDigits: "1234",
			},
			expectError: true,
		},
		{
			name: "Invalid - Negative Amount",
			invoice: &InvoiceDomain{
				ID:             uuid.New(),
				AccountID:      uuid.New(),
				Amount:         -10.0,
				Status:         "pending",
				Description:    "Test invoice",
				CardLastDigits: "1234",
			},
			expectError: true,
		},
		{
			name: "Invalid - Empty Description",
			invoice: &InvoiceDomain{
				ID:             uuid.New(),
				AccountID:      uuid.New(),
				Amount:         100.0,
				Status:         "pending",
				Description:    "",
				CardLastDigits: "1234",
			},
			expectError: true,
		},
		{
			name: "Invalid - Empty CardLastDigits",
			invoice: &InvoiceDomain{
				ID:             uuid.New(),
				AccountID:      uuid.New(),
				Amount:         100.0,
				Status:         "pending",
				Description:    "Test invoice",
				CardLastDigits: "",
			},
			expectError: true,
		},
		{
			name: "Invalid - Invalid Status",
			invoice: &InvoiceDomain{
				ID:             uuid.New(),
				AccountID:      uuid.New(),
				Amount:         100.0,
				Status:         "invalid_status",
				Description:    "Test invoice",
				CardLastDigits: "1234",
			},
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.invoice.IsValid()
			if test.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
