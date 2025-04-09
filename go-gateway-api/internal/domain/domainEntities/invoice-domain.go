package domainEntities

import (
	"math/rand"
	"time"

	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/domain/valueObjects"
	"github.com/google/uuid"
)

// InvoiceDomain represents the invoice domain entity
// Status represents the invoice status
type Status string

const (
	StatusPending  Status = "pending"
	StatusApproved Status = "approved"
	StatusFailed   Status = "failed"
)

// InvoiceDomain represents the invoice domain entity
type InvoiceDomain struct {
	ID             string
	AccountID      string
	Amount         float64
	Status         Status
	Description    string
	PaymentType    string
	CardLastDigits string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// NewInvoice creates a new invoice with default values
func NewInvoice(accountID string, amount float64, description string, paymentType string, card *valueObjects.CreditCard) (*InvoiceDomain, error) {
	if amount <= 0 {
		return nil, ErrInvalidAmount
	}

	cardLastDigits := ""
	if card != nil {
		cardLastDigits = card.GetLastDigits()
	}

	invoice := &InvoiceDomain{
		ID:             uuid.New().String(),
		AccountID:      accountID,
		Amount:         amount,
		Status:         StatusPending,
		Description:    description,
		PaymentType:    paymentType,
		CardLastDigits: cardLastDigits,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	return invoice, nil
}

// SetStatus updates the status of the invoice and the updated_at timestamp
func (i *InvoiceDomain) SetStatus(status Status) {
	if i.Status != StatusPending {
		panic(ErrInvalidStatus)
	}
	i.Status = status
	i.UpdatedAt = time.Now()
}

func (i *InvoiceDomain) Process() error {
	if i.Amount <= 0 {
		return ErrInvalidAmount
	}
	randomSource := rand.New(rand.NewSource(time.Now().UnixNano()))
	if randomSource.Float64() < 0.7 {
		i.SetStatus(StatusApproved)
	} else {
		i.SetStatus(StatusFailed)
	}
	i.Status = StatusApproved
	i.UpdatedAt = time.Now()
	return nil
}
