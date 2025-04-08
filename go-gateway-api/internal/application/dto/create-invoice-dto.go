package dto

type CreateInvoiceInputDTO struct {
	APIKey          string
	AccountID       string  `json:"account_id" validate:"required"`
	Amount          float64 `json:"amount" validate:"required"`
	Description     string  `json:"description" validate:"required"`
	PaymentType     string  `json:"payment_type" validate:"required"`
	CardNumber      string  `json:"card_number" validate:"required"`
	CVV             string  `json:"cvv" validate:"required"`
	ExpirationMonth int     `json:"expiration_month" validate:"required"`
	ExpirationYear  int     `json:"expiration_year" validate:"required"`
	CardHolderName  string  `json:"card_holder_name" validate:"required"`
}
