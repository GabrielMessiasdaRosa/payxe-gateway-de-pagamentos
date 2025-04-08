package dto

type CreateInvoiceInputDTO struct {
	Amount      float64 `json:"amount" validate:"required"`
	Description string  `json:"description" validate:"required"`
	AccountID   string  `json:"account_id" validate:"required"`
}
