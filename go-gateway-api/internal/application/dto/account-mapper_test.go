package dto

import (
	"testing"

	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/domain/domainEntities"
	"github.com/stretchr/testify/assert"
)

func TestFromAccount(t *testing.T) {
	t.Run("Should convert AccountDomain to AccountOutputDTO", func(t *testing.T) {
		// Arrange
		accountDomain := domainEntities.NewAccount("Test Account", "test@example.com")
		accountDomain.ID = "test-id"
		accountDomain.Balance = 100.50
		accountDomain.APIKey = "test-api-key"

		// Act
		result := FromAccount(accountDomain)

		// Assert
		assert.Equal(t, accountDomain.ID, result.ID)
		assert.Equal(t, accountDomain.Name, result.Name)
		assert.Equal(t, accountDomain.Email, result.Email)
		assert.Equal(t, accountDomain.Balance, result.Balance)
		assert.Equal(t, accountDomain.APIKey, result.APIKey)
	})

	t.Run("Should handle zero values correctly", func(t *testing.T) {
		// Arrange
		accountDomain := domainEntities.NewAccount("", "")
		accountDomain.ID = ""
		accountDomain.Balance = 0
		accountDomain.APIKey = ""

		// Act
		result := FromAccount(accountDomain)

		// Assert
		assert.Equal(t, "", result.ID)
		assert.Equal(t, "", result.Name)
		assert.Equal(t, "", result.Email)
		assert.Equal(t, 0.0, result.Balance)
		assert.Equal(t, "", result.APIKey)
	})
}