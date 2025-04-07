package service

import (
	"errors"
	"testing"

	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/application/dto"
	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/infra/mockRepositories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAccountService_Save(t *testing.T) {
	t.Run("should save account successfully", func(t *testing.T) {
		// Arrange
		mockRepo := new(mockRepositories.MockAccountRepository)
		service := NewAccountService(mockRepo)
		
		input := &dto.CreateAccountInputDTO{
			Name:        "Test Account",
			Email:       "test@example.com",

		}
		
		mockRepo.On("Save", mock.Anything).Return(nil)
		
		// Act
		output, err := service.Save(input)
		
		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, output)
		assert.Equal(t, input.Name, output.Name)
		assert.Equal(t, input.Email, output.Email)
		mockRepo.AssertExpectations(t)
	})
	t.Run("should return error when repository fails", func(t *testing.T) {
		// Arrange
		mockRepo := new(mockRepositories.MockAccountRepository)
		service := NewAccountService(mockRepo)
		
		input := &dto.CreateAccountInputDTO{
			Name:        "Test Account",
			Email:       "test@example.com",
		}
		
		expectedError := errors.New("repository error")
		mockRepo.On("Save", mock.Anything).Return(expectedError)
		mockRepo.On("Save", mock.Anything).Return(expectedError)
		
		// Act
		output, err := service.Save(input)
		
		// Assert
		assert.Error(t, err)
		assert.Nil(t, output)
		mockRepo.AssertExpectations(t)
	})
}

func TestAccountService_FindByAPIKey(t *testing.T) {
	t.Run("should find account by API key successfully", func(t *testing.T) {

		mockRepo := new(mockRepositories.MockAccountRepository)
		service := NewAccountService(mockRepo)
		
		apiKey := "test-api-key"
		expectedAccount := &dto.AccountOutputDTO{
			Name:  "Test Account",
			Email: "test@example.com",
		}
		

		mockRepo.On("FindByAPIKey", apiKey).Return(expectedAccount, nil)
		// Act
		account, err := service.accountRepository.FindByAPIKey(apiKey)
		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, account)
		assert.Equal(t, expectedAccount.Name, account.Name)
		assert.Equal(t, expectedAccount.Email, account.Email)
		mockRepo.AssertExpectations(t)
})}
