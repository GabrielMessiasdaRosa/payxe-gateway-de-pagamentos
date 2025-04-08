package service_test

import (
	"testing"

	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/application/dto"
	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/application/service"
	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/infra/mockRepositories"
)

func TestAccountService(t *testing.T) {
	// @audit-issue covarage na linha 23
	t.Run("should create an account successfully", func(t *testing.T) {
		repo := mockRepositories.NewInMemoryAccountRepository()
		accService := service.NewAccountService(repo)

		input := &dto.CreateAccountInputDTO{
			Name:  "John Doe",
			Email: "john@example.com",
		}

		output, err := accService.CreateAccount(input)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if output.Name != input.Name {
			t.Errorf("expected name %s, got %s", input.Name, output.Name)
		}
		if output.Email != input.Email {
			t.Errorf("expected email %s, got %s", input.Email, output.Email)
		}
		if output.ID == "" {
			t.Errorf("expected non-empty ID")
		}
		if output.APIKey == "" {
			t.Errorf("expected non-empty APIKey")
		}
	})
	// @audit-ok
	t.Run("Should return an account when findByAPIKey is called", func(t *testing.T) {
		repo := mockRepositories.NewInMemoryAccountRepository()
		accService := service.NewAccountService(repo)

		// Create an account first.
		input := &dto.CreateAccountInputDTO{
			Name:  "Alice",
			Email: "alice@example.com",
		}
		created, err := accService.CreateAccount(input)
		if err != nil {
			t.Fatalf("error creating account: %v", err)
		}

		// Successful lookup.
		found, err := accService.FindByAPIKey(created.APIKey)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if found.ID != created.ID {
			t.Errorf("expected ID %s, got %s", created.ID, found.ID)
		}

		// Negative lookup: non-existent APIKey.
		_, err = accService.FindByAPIKey("non-existent-apikey")
		if err == nil {
			t.Errorf("expected error for non-existent APIKey")
		}
	})
	// @audit-ok
	t.Run("Should return an account when findByID is called", func(t *testing.T) {
		repo := mockRepositories.NewInMemoryAccountRepository()
		accService := service.NewAccountService(repo)

		// Create an account first.
		input := &dto.CreateAccountInputDTO{
			Name:  "Bob",
			Email: "bob@example.com",
		}
		created, err := accService.CreateAccount(input)
		if err != nil {
			t.Fatalf("error creating account: %v", err)
		}

		// Successful lookup.
		found, err := accService.FindByID(created.ID)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if found.APIKey != created.APIKey {
			t.Errorf("expected APIKey %s, got %s", created.APIKey, found.APIKey)
		}

		// Negative lookup: non-existent ID.
		_, err = accService.FindByID("non-existent-id")
		if err == nil {
			t.Errorf("expected error for non-existent ID")
		}
	})

}

// @audit-ok
func TestUpdateBalance(t *testing.T) {
	repo := mockRepositories.NewInMemoryAccountRepository()
	accService := service.NewAccountService(repo)

	// Create an account first.
	input := &dto.CreateAccountInputDTO{
		Name:  "Charlie",
		Email: "charlie@example.com",
	}
	created, err := accService.CreateAccount(input)
	if err != nil {
		t.Fatalf("error creating account: %v", err)
	}

	// Successful balance update.
	updateInput := &dto.UpdateAccountInputDTO{
		ID: created.ID,
	}
	err = accService.UpdateBalance(updateInput)
	if err != nil {
		t.Fatalf("expected no error updating balance, got %v", err)
	}

	// Negative: update balance for a non-existent account.
	updateInput.ID = "non-existent-id"
	err = accService.UpdateBalance(updateInput)
	if err == nil {
		t.Errorf("expected error when updating balance for non-existent account")
	}

	// @audit-ok
	t.Run("should return error when account not found", func(t *testing.T) {
		// Arrange
		repo := mockRepositories.NewInMemoryAccountRepository()
		accService := service.NewAccountService(repo)

		input := &dto.UpdateAccountInputDTO{
			ID:      "non-existent-id",
			Balance: 100.0,
		}

		// Act
		err := accService.UpdateBalance(input)

		// Assert
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	// @audit-ok
	t.Run("should update balance from 100 to 200 successfully", func(t *testing.T) {
		// Arrange
		repo := mockRepositories.NewInMemoryAccountRepository()
		accService := service.NewAccountService(repo)

		createInput := &dto.CreateAccountInputDTO{
			Name:  "Balance Test",
			Email: "balance@example.com",
		}
		created, err := accService.CreateAccount(createInput)
		if err != nil {
			t.Fatalf("error creating account: %v", err)
		}

		// Set initial balance to 100
		updateInput := &dto.UpdateAccountInputDTO{
			ID:      created.ID,
			Balance: 100.0,
		}
		err = accService.UpdateBalance(updateInput)
		if err != nil {
			t.Fatalf("failed to set initial balance: %v", err)
		}

		// Act - Update balance to 200
		updateInput.Balance = 200.0
		err = accService.UpdateBalance(updateInput)

		// Assert
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		// Verify the balance was updated
		account, err := accService.FindByID(created.ID)
		if err != nil {
			t.Fatalf("failed to retrieve account: %v", err)
		}

		if account.Balance != 200.0 {
			t.Errorf("expected balance 200.0, got %f", account.Balance)
		}
	})
}
