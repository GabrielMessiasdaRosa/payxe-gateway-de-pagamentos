package service

import (
	"errors"
	"testing"

	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/application/dto"
	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/infra/mockRepositories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAccountService_CreateAccount(t *testing.T) {
	t.Run("should create account successfully", func(t *testing.T) {
		// Arrange
		mockRepo := new(mockRepositories.MockAccountRepository)
		service := NewAccountService(mockRepo)

		input := &dto.CreateAccountInputDTO{
			Name:  "Test Account",
			Email: "test@example.com",
		}

		mockRepo.On("CreateAccount", mock.Anything).Return(nil)

		// Act
		output, err := service.CreateAccount(input)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, output)
		assert.NotNil(t, output.ID)
		assert.NotEmpty(t, output.APIKey)
		assert.Equal(t, input.Name, output.Name)
		assert.Equal(t, input.Email, output.Email)
		mockRepo.AssertExpectations(t)
	})
	t.Run("should return error when repository fails", func(t *testing.T) {
		// Arrange
		mockRepo := new(mockRepositories.MockAccountRepository)
		service := NewAccountService(mockRepo)

		input := &dto.CreateAccountInputDTO{
			Name:  "Test Account",
			Email: "test@example.com",
		}

		expectedError := errors.New("repository error")
		mockRepo.On("CreateAccount", mock.Anything).Return(expectedError)
		mockRepo.On("CreateAccount", mock.Anything).Return(expectedError)

		// Act
		output, err := service.CreateAccount(input)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, output)
		mockRepo.AssertExpectations(t)
	})
}

func TestAccountService_FindByAPIKey(t *testing.T) {
	t.Run("should find account by API key successfully", func(t *testing.T) {
		// Arrange
		mockRepo := new(mockRepositories.MockAccountRepository)
		service := NewAccountService(mockRepo)

		input := &dto.CreateAccountInputDTO{
			Name:  "Test Account",
			Email: "test@example.com",
		}

		mockRepo.On("CreateAccount", mock.Anything).Return(nil)
		// Act
		output, err := service.CreateAccount(input)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, output)
		assert.NotNil(t, output.ID)
		assert.NotEmpty(t, output.APIKey)
		assert.Equal(t, input.Name, output.Name)
		assert.Equal(t, input.Email, output.Email)
		mockRepo.AssertExpectations(t)

		mockRepo.On("FindByAPIKey", output.APIKey).Return(output, nil)
		account, err := service.FindByAPIKey(output.APIKey)
		assert.NoError(t, err)
		assert.NotNil(t, account)
		assert.Equal(t, output.ID, account.ID)
		assert.Equal(t, output.Name, account.Name)
		assert.Equal(t, input.Email, account.Email)
		assert.Equal(t, output.Balance, account.Balance)
		assert.Equal(t, output.APIKey, account.APIKey)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when account not found", func(t *testing.T) {
		// Arrange
		mockRepo := new(mockRepositories.MockAccountRepository)
		service := NewAccountService(mockRepo)

		apiKey := "non-existent-api-key"
		expectedError := errors.New("account not found")
		mockRepo.On("FindByAPIKey", apiKey).Return(nil, expectedError)

		// Act
		account, err := service.FindByAPIKey(apiKey)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, account)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return account by ID successfully", func(t *testing.T) {

	})

	t.Run("should update account balance successfully", func(t *testing.T) {
		// Arrange
		mockRepo := new(mockRepositories.MockAccountRepository)
		service := NewAccountService(mockRepo)

		input := &dto.CreateAccountInputDTO{
			Name:  "Test Account",
			Email: "test@example.com",
		}

		mockRepo.On("CreateAccount", mock.Anything).Return(nil)

		// Act
		output, _ := service.CreateAccount(input)

		// Arrange
		updateInput := &dto.UpdateAccountInputDTO{
			ID:      output.ID,
			Balance: 100.0,
			Name:    "Updated Account",
			Email:   "updated@example.com",
			APIKey:  "updated-api-key",
		}

		mockRepo.On("FindByID", output.ID).Return(output, nil)
		mockRepo.On("UpdateBalance", output).Return(nil)
		// Act
		err := service.UpdateBalance(updateInput)
		// Assert
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)

	})
}
