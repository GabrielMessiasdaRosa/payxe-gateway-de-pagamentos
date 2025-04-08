package dto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToAccountDomain(t *testing.T) {
	t.Run("should convert AccountOutputDTO to AccountDomain", func(t *testing.T) {
		// Arrange
		input := AccountOutputDTO{
			ID:      "account-id-123",
			Name:    "Test User",
			Email:   "test@example.com",
			Balance: 150.75,
			APIKey:  "api-key-456",
		}

		// Act
		result := ToAccountDomain(input)

		// Assert
		assert.Equal(t, input.Name, result.Name)
		assert.Equal(t, input.Email, result.Email)
		assert.NotEmpty(t, result.ID)
		assert.NotEmpty(t, result.APIKey)

	})

	t.Run("should convert CreateAccountInputDTO to AccountDomain", func(t *testing.T) {
		// Arrange
		input := CreateAccountInputDTO{
			Name:  "Test User",
			Email: "test@example.com",
		}

		// Act
		result := ToAccountDomain(input)

		// Assert
		assert.Equal(t, result.ID, "")
		assert.Equal(t, input.Name, result.Name)
		assert.Equal(t, input.Email, result.Email)
		assert.Equal(t, result.Balance, 0.0)
		assert.Equal(t, result.APIKey, "")
	})

	t.Run("should convert UpdateAccountInputDTO to AccountDomain", func(t *testing.T) {
		// Arrange

		input := UpdateAccountInputDTO{
			ID:      "account-id-123",
			Balance: 100.50,
			Name:    "Test User",
			Email:   "test@example.com",
			APIKey:  "api-key-456",
		}

		// Act
		result := ToAccountDomain(input)

		// Assert
		assert.Equal(t, input.ID, result.ID)
		assert.Equal(t, input.Balance, result.Balance)
		assert.Equal(t, input.Name, result.Name)
		assert.Equal(t, input.Email, result.Email)
		assert.Equal(t, input.APIKey, result.APIKey)
	})
	t.Run("should handle nil input", func(t *testing.T) {
		// Act
		result := ToAccountDomain(nil)

		// Assert
		assert.Nil(t, result)
	})
	t.Run("should handle unsupported type", func(t *testing.T) {
		// Arrange
		type UnsupportedType struct {
			Field string
		}
		input := UnsupportedType{Field: "test"}

		// Act
		result := ToAccountDomain(input)

		// Assert
		assert.Nil(t, result)
	})
}
