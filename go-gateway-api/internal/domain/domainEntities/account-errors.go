package domainEntities

import "errors"

var (
	ErrAccountNotFound       = errors.New("account not found")
	ErrAccountDuplicateAPIKey = errors.New("account already has an API key")
	ErrInvoiceNotFound		= errors.New("invoice not found")
	ErrUnauthorizedAccess	= errors.New("unauthorized access")
)