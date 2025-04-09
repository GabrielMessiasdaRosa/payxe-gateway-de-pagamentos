package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/application/dto"
	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/application/service"
	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/infra/handlers"
	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/infra/mockRepositories"
	"github.com/stretchr/testify/assert"
)

func setupInvoiceService() *service.InvoiceService {
	accountRepo := mockRepositories.NewInMemoryAccountRepository() // Usar repositório mock
	accountService := service.NewAccountService(accountRepo)
	invoiceRepo := mockRepositories.NewInMemoryInvoiceRepository()
	invoiceService := service.NewInvoiceService(invoiceRepo, accountService)
	return invoiceService
}

func TestCreateInvoice_Success(t *testing.T) {
	invoiceService := setupInvoiceService()
	handler := handlers.NewInvoiceHandler(invoiceService)

	// Create an account to associate with the invoice
	accountInput := &dto.CreateAccountInputDTO{
		Name:  "Test Account",
		Email: "test@example.com",
	}
	account, _ := invoiceService.AccountService.CreateAccount(accountInput)

	input := dto.CreateInvoiceInputDTO{
		AccountID:       account.ID,
		Amount:          100.50,
		Description:     "Test Invoice",
		PaymentType:     "credit_card",
		CardNumber:      "4111111111111111",
		CVV:             "123",
		ExpirationMonth: 12,
		ExpirationYear:  2030,
		CardHolderName:  "John Doe",
	}
	body, _ := json.Marshal(input)
	req, _ := http.NewRequest("POST", "/invoices", bytes.NewBuffer(body))
	req.Header.Set("X-API-Key", account.APIKey)
	rr := httptest.NewRecorder()

	handler.Create(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var response dto.InvoiceOutputDTO
	json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NotEmpty(t, response.ID)
	assert.Equal(t, input.AccountID, response.AccountID)
	assert.Equal(t, input.Amount, response.Amount)
	assert.Equal(t, input.Description, response.Description)
}

func TestCreateInvoice_InvalidInput(t *testing.T) {
	invoiceService := setupInvoiceService()
	handler := handlers.InvoiceHandler{InvoiceService: invoiceService}

	// Invalid JSON
	req, _ := http.NewRequest("POST", "/invoices", bytes.NewBuffer([]byte("invalid json")))
	rr := httptest.NewRecorder()

	handler.Create(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Contains(t, rr.Body.String(), "Invalid JSON")
}

func TestCreateInvoice_MissingFields(t *testing.T) {
	invoiceService := setupInvoiceService()
	handler := handlers.InvoiceHandler{InvoiceService: invoiceService}

	// Missing fields
	input := dto.CreateInvoiceInputDTO{
		AccountID:   "",
		Amount:      0,
		Description: "",
	}
	body, _ := json.Marshal(input)
	req, _ := http.NewRequest("POST", "/invoices", bytes.NewBuffer(body))
	req.Header.Set("X-API-Key", "some-api-key")
	rr := httptest.NewRecorder()

	handler.Create(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Contains(t, rr.Body.String(), "Missing or invalid fields")
}

func TestCreateInvoice_InvalidAPIKey(t *testing.T) {
	invoiceService := setupInvoiceService()
	handler := handlers.InvoiceHandler{InvoiceService: invoiceService}

	// Valid input but invalid API key
	input := dto.CreateInvoiceInputDTO{
		AccountID:   "some-account-id",
		Amount:      100.50,
		Description: "Test Invoice",
	}
	body, _ := json.Marshal(input)
	req, _ := http.NewRequest("POST", "/invoices", bytes.NewBuffer(body))
	req.Header.Set("X-API-Key", "invalid-api-key")
	rr := httptest.NewRecorder()

	handler.Create(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Contains(t, rr.Body.String(), "account not found")
}
