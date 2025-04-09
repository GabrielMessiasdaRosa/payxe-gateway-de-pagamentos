package mockRepositories

import (
	"testing"

	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/domain/domainEntities"
	"github.com/stretchr/testify/assert"
)

func TestNewInMemoryAccountRepository(t *testing.T) {
	repo := NewInMemoryAccountRepository()
	assert.NotNil(t, repo)
	assert.Empty(t, repo.accounts)
}

func TestCreateAccount(t *testing.T) {
	repo := NewInMemoryAccountRepository()
	account := &domainEntities.AccountDomain{
		ID:      "1",
		APIKey:  "api-key-1",
		Balance: 100,
	}

	err := repo.CreateAccount(account)
	assert.Nil(t, err)
	assert.Len(t, repo.accounts, 1)
	assert.Equal(t, account, repo.accounts[0])
}

func TestFindByID(t *testing.T) {
	repo := NewInMemoryAccountRepository()
	account1 := &domainEntities.AccountDomain{
		ID:      "1",
		APIKey:  "api-key-1",
		Balance: 100,
	}
	account2 := &domainEntities.AccountDomain{
		ID:      "2",
		APIKey:  "api-key-2",
		Balance: 200,
	}

	repo.CreateAccount(account1)
	repo.CreateAccount(account2)

	t.Run("should find account by ID", func(t *testing.T) {
		found, err := repo.FindByID("1")
		assert.Nil(t, err)
		assert.Equal(t, account1, found)

		found, err = repo.FindByID("2")
		assert.Nil(t, err)
		assert.Equal(t, account2, found)
	})

	t.Run("should return error when account not found", func(t *testing.T) {
		found, err := repo.FindByID("non-existent")
		assert.Error(t, err)
		assert.Equal(t, "account not found", err.Error())
		assert.Nil(t, found)
	})
}

func TestFindByAPIKey(t *testing.T) {
	repo := NewInMemoryAccountRepository()
	account1 := &domainEntities.AccountDomain{
		ID:      "1",
		APIKey:  "api-key-1",
		Balance: 100,
	}
	account2 := &domainEntities.AccountDomain{
		ID:      "2",
		APIKey:  "api-key-2",
		Balance: 200,
	}

	repo.CreateAccount(account1)
	repo.CreateAccount(account2)

	t.Run("should find account by API key", func(t *testing.T) {
		found, err := repo.FindByAPIKey("api-key-1")
		assert.Nil(t, err)
		assert.Equal(t, account1, found)

		found, err = repo.FindByAPIKey("api-key-2")
		assert.Nil(t, err)
		assert.Equal(t, account2, found)
	})

	t.Run("should return error when account not found", func(t *testing.T) {
		found, err := repo.FindByAPIKey("non-existent")
		assert.Error(t, err)
		assert.Equal(t, "account not found", err.Error())
		assert.Nil(t, found)
	})
}

func TestUpdateBalance(t *testing.T) {
	repo := NewInMemoryAccountRepository()
	account := &domainEntities.AccountDomain{
		ID:      "1",
		APIKey:  "api-key-1",
		Balance: 100,
	}

	repo.CreateAccount(account)

	t.Run("should update account balance", func(t *testing.T) {
		updatedAccount := &domainEntities.AccountDomain{
			ID:      "1",
			APIKey:  "api-key-1",
			Balance: 200,
		}

		err := repo.UpdateBalance(updatedAccount)
		assert.Nil(t, err)

		found, _ := repo.FindByID("1")
		assert.Equal(t, float64(200), found.Balance)
	})

	t.Run("should return error when account not found", func(t *testing.T) {
		nonExistentAccount := &domainEntities.AccountDomain{
			ID:      "non-existent",
			Balance: 300,
		}

		err := repo.UpdateBalance(nonExistentAccount)
		assert.Error(t, err)
		assert.Equal(t, "account not found", err.Error())
	})
}
