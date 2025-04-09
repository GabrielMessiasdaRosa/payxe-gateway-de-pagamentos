package dto

type UpdateInvoiceInputDTO struct {
	ID          string  `json:"id" validate:"required"`
	Amount      float64 `json:"amount" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Status      string  `json:"status" validate:"required"`
	AccountID   string  `json:"account_id" validate:"required"`
}
