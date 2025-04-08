package domainEntities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// InvoiceDomain represents the invoice domain entity
// Status represents the invoice status
type Status string

const (
	StatusPending Status = "pending"
	StatusPaid    Status = "paid"
	StatusFailed  Status = "failed"
)

// InvoiceDomain represents the invoice domain entity
type InvoiceDomain struct {
	ID             uuid.UUID
	AccountID      uuid.UUID
	Amount         float64
	Status         Status
	Description    string
	PaymentType    string
	CardLastDigits string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// NewInvoice creates a new invoice with default values
func NewInvoice(accountID uuid.UUID, amount float64, description, cardLastDigits string) *InvoiceDomain {
	now := time.Now()
	return &InvoiceDomain{
		ID:             uuid.New(),
		AccountID:      accountID,
		Amount:         amount,
		Status:         "pending",
		Description:    description,
		CardLastDigits: cardLastDigits,
		CreatedAt:      now,
		UpdatedAt:      now,
	}
}

// SetStatus updates the status of the invoice and the updated_at timestamp
func (i *InvoiceDomain) SetStatus(status Status) {
	i.Status = status
	i.UpdatedAt = time.Now()
}

// IsValid validates the invoice fields
func (i *InvoiceDomain) IsValid() error {
	if i.AccountID == uuid.Nil {
		return errors.New("account ID cannot be empty")
	}

	if i.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	if i.Description == "" {
		return errors.New("description cannot be empty")
	}

	if i.CardLastDigits == "" {
		return errors.New("card last digits cannot be empty")
	}

	// Validate status
	validStatuses := map[Status]bool{
		StatusPending: true,
		StatusPaid:    true,
		StatusFailed:  true,
	}

	if !validStatuses[i.Status] {
		return errors.New("invalid status")
	}

	return nil
}
