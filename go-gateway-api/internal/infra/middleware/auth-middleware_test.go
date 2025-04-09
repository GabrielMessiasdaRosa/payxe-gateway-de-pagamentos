package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/application/service"
	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/domain/domainEntities"
	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/infra/middleware"
	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/infra/mockRepositories"
	"github.com/stretchr/testify/assert"
)

func setupAccountService() *service.AccountService {
	// Configure o AccountService real aqui.
	// Por exemplo, você pode usar um banco de dados em memória ou um mock de repositório.
	accountRepo := mockRepositories.NewInMemoryAccountRepository() // Exemplo de repositório em memória
	accountService := service.NewAccountService(accountRepo)

	// Adicione dados de teste ao repositório
	accountRepo.CreateAccount(&domainEntities.AccountDomain{
		APIKey: "valid-api-key",
	})

	return accountService
}

func TestAuthMiddleware_Authenticate_Success(t *testing.T) {
	accountService := setupAccountService()
	authMiddleware := middleware.NewAuthMiddleware(accountService)

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	handler := authMiddleware.Authenticate(nextHandler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("X-API-Key", "valid-api-key")
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestAuthMiddleware_Authenticate_MissingAPIKey(t *testing.T) {
	accountService := setupAccountService()
	authMiddleware := middleware.NewAuthMiddleware(accountService)

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	handler := authMiddleware.Authenticate(nextHandler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.Equal(t, "API key is required\n", rec.Body.String())
}

func TestAuthMiddleware_Authenticate_InvalidAPIKey(t *testing.T) {
	accountService := setupAccountService()
	authMiddleware := middleware.NewAuthMiddleware(accountService)

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	handler := authMiddleware.Authenticate(nextHandler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("X-API-Key", "invalid-api-key")
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.Equal(t, "Invalid API key\n", rec.Body.String())
}

