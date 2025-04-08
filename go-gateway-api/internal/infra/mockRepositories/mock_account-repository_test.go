package mockRepositories

import (
	"errors"
	"testing"

	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/domain/domainEntities"
	"github.com/stretchr/testify/assert"
)

func TestMockAccountRepository_FindByAPIKey(t *testing.T) {
	t.Run("should return account when FindByAPIKey is called", func(t *testing.T) {
		// Arrange
		mockRepo := new(MockAccountRepository)
		apiKey := "test-api-key"
		expectedAccount := domainEntities.AccountDomain{
			ID:    "acc123",
			Name:  "Test Account",
			Email: "test@example.com",
		}

		mockRepo.On("FindByAPIKey", apiKey).Return(&expectedAccount, nil)

		// Act
		account, err := mockRepo.FindByAPIKey(apiKey)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, account)
		assert.Equal(t, &expectedAccount, account)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when FindByAPIKey fails", func(t *testing.T) {
		// Arrange
		mockRepo := new(MockAccountRepository)
		apiKey := "invalid-api-key"
		expectedError := errors.New("account not found")

		mockRepo.On("FindByAPIKey", apiKey).Return(nil, expectedError)

		// Act
		account, err := mockRepo.FindByAPIKey(apiKey)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, account)
		assert.Equal(t, expectedError, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should handle nil account return", func(t *testing.T) {
		// Arrange
		mockRepo := new(MockAccountRepository)
		apiKey := "api-key-with-no-account"

		mockRepo.On("FindByAPIKey", apiKey).Return(nil, nil)

		// Act
		account, err := mockRepo.FindByAPIKey(apiKey)

		// Assert
		assert.NoError(t, err)
		assert.Nil(t, account)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Should update balance successfully", func(t *testing.T) {
		// Arrange
		mockRepo := new(MockAccountRepository)
		expectedAccount := &domainEntities.AccountDomain{
			ID:      "acc123",
			Name:    "Test Account",
			Email:   "test@example.com",
			Balance: 100.0,
		}
		newBalance := 150.0
		expectedAccount.Balance = newBalance
		mockRepo.On("UpdateBalance", expectedAccount).Return(nil)
		mockRepo.On("FindByID", expectedAccount.ID).Return(expectedAccount, nil)

		// Act
		mockRepo.UpdateBalance(expectedAccount)
		updatedAcc, err := mockRepo.FindByID(expectedAccount.ID)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, updatedAcc)
		assert.Equal(t, newBalance, updatedAcc.Balance)
		mockRepo.AssertExpectations(t)
	})
}
