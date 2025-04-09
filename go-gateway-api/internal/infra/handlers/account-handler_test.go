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

func setupAccountService() *service.AccountService {
	// Configurar o repositório real (pode ser um repositório em memória ou um banco de dados real)
	accountRepo := mockRepositories.NewInMemoryAccountRepository() // Usar repositório mock
	return service.NewAccountService(accountRepo)
}

// @audit-ok
func TestNewAccountHandler(t *testing.T) {
	accService := setupAccountService()
	handler := handlers.NewAccountHandler(accService)
	assert.NotNil(t, handler)
	assert.Equal(t, accService, handler.AccountService)
	assert.IsType(t, &handlers.AccountHandler{}, handler)
}

// @audit-issue coverage na linha 30~31
func TestCreate_Success(t *testing.T) {
	accService := setupAccountService()
	handler := handlers.AccountHandler{AccountService: accService}

	input := &dto.CreateAccountInputDTO{
		Name:  "Test Account",
		Email: "test@example.com",
	}

	body, _ := json.Marshal(input)
	req, _ := http.NewRequest("POST", "/accounts", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	handler.Create(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var response dto.AccountOutputDTO
	json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NotEmpty(t, response.ID)
	assert.Equal(t, input.Name, response.Name)
	assert.Equal(t, input.Email, response.Email)
}

func TestCreate_InvalidInput(t *testing.T) {
	accService := setupAccountService()
	handler := handlers.AccountHandler{AccountService: accService}

	// Invalid JSON
	req, _ := http.NewRequest("POST", "/accounts", bytes.NewBuffer([]byte("invalid json")))
	rr := httptest.NewRecorder()

	handler.Create(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)

}

// @audit-issue coverage na linha 47~48
func TestGet_Success(t *testing.T) {
	accService := setupAccountService()
	handler := handlers.AccountHandler{AccountService: accService}

	// Criar uma conta para teste
	account, _ := accService.CreateAccount(&dto.CreateAccountInputDTO{
		Name:  "Test Account",
		Email: "test@example.com",
	})

	req, _ := http.NewRequest("GET", "/accounts", nil)
	req.Header.Set("X-API-Key", account.APIKey) // Usar ID como API Key para o exemplo
	rr := httptest.NewRecorder()

	handler.Get(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response dto.AccountOutputDTO
	json.Unmarshal(rr.Body.Bytes(), &response)
	assert.Equal(t, account.ID, response.ID)
	assert.Equal(t, account.Name, response.Name)
	assert.Equal(t, account.Email, response.Email)
}

func TestGet_NoAPIKey(t *testing.T) {
	accService := setupAccountService()
	handler := handlers.AccountHandler{AccountService: accService}

	req, _ := http.NewRequest("GET", "/accounts", nil)
	rr := httptest.NewRecorder()

	handler.Get(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Contains(t, rr.Body.String(), "API key is required")
}

func TestGet_AccountNotFound(t *testing.T) {
	accService := setupAccountService()
	handler := handlers.AccountHandler{AccountService: accService}

	req, _ := http.NewRequest("GET", "/accounts", nil)
	req.Header.Set("X-API-Key", "non-existent-api-key")
	rr := httptest.NewRecorder()

	handler.Get(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Contains(t, rr.Body.String(), "account not found")

}
