package controllers_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/application/dto"
	controllers "github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/infra/handlers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAccountService struct {
	mock.Mock
}

func (m *MockAccountService) CreateAccount(input *dto.CreateAccountInputDTO) (*dto.AccountOutputDTO, error) {
	args := m.Called(input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.AccountOutputDTO), args.Error(1)
}

func (m *MockAccountService) FindByAPIKey(apiKey string) (*dto.AccountOutputDTO, error) {
	args := m.Called(apiKey)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.AccountOutputDTO), args.Error(1)
}

func TestAccountHandler_FindByAPIKey(t *testing.T) {
	t.Run("should return account when API key exists", func(t *testing.T) {
		mockService := new(MockAccountService)
		handler := controllers.NewAccountHandler(mockService)

		expectedAccount := &dto.AccountOutputDTO{
			ID:      "1",
			Name:    "Test User",
			Email:   "test@example.com",
			Balance: 100.0,
			APIKey:  "test-api-key",
		}

		mockService.On("FindByAPIKey", "test-api-key").Return(expectedAccount, nil)

		req, err := http.NewRequest("GET", "/accounts?apiKey=test-api-key", nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		handler.FindByAPIKey(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var responseAccount dto.AccountOutputDTO
		err = json.Unmarshal(rr.Body.Bytes(), &responseAccount)
		assert.NoError(t, err)
		assert.Equal(t, expectedAccount, &responseAccount)
		mockService.AssertExpectations(t)
	})

	t.Run("should return error when API key is not provided", func(t *testing.T) {
		mockService := new(MockAccountService)
		handler := controllers.NewAccountHandler(mockService)

		req, err := http.NewRequest("GET", "/accounts", nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		handler.FindByAPIKey(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Contains(t, rr.Body.String(), "API key is required")
		mockService.AssertNotCalled(t, "FindByAPIKey")
	})

	t.Run("should return error when service returns error", func(t *testing.T) {
		mockService := new(MockAccountService)
		handler := controllers.NewAccountHandler(mockService)

		mockService.On("FindByAPIKey", "invalid-api-key").Return(nil, errors.New("account not found"))

		req, err := http.NewRequest("GET", "/accounts?apiKey=invalid-api-key", nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		handler.FindByAPIKey(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Contains(t, rr.Body.String(), "account not found")
		mockService.AssertExpectations(t)
	})
}
