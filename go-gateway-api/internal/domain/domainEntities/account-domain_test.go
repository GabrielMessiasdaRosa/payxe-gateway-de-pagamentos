package domainEntities

import (
	"sync"
	"testing"
)

// @audit-ok // deve criar uma conta com os dados corretos
func TestNewAccount(t *testing.T) {
	name := "John Doe"
	email := "john.doe@example.com"

	account := NewAccount(name, email)
	if account.Name != name {
		t.Errorf("expected Name to be %s, got %s", name, account.Name)
	}

	if account.Email != email {
		t.Errorf("expected Email to be %s, got %s", email, account.Email)
	}

	if account.Balance != 0 {
		t.Errorf("expected Balance to be 0, got %f", account.Balance)
	}

	if account.APIKey == "" {
		t.Error("expected APIKey to be generated, but it is empty")
	}

	if account.CreatedAt.IsZero() {
		t.Error("expected CreatedAt to be set, but it is zero")
	}

	if account.UpdatedAt.IsZero() {
		t.Error("expected UpdatedAt to be set, but it is zero")
	}
}

// @audit-ok // deve adicionar saldo à conta e atualizar o UpdatedAt
func TestAddBalance(t *testing.T) {
	account := NewAccount("Jane Doe", "jane.doe@example.com")
	initialBalance := account.Balance
	amount := 100.0

	account.AddBalance(amount)

	if account.Balance != initialBalance+amount {
		t.Errorf("expected Balance to be %f, got %f", initialBalance+amount, account.Balance)
	}

	if account.UpdatedAt.Before(account.CreatedAt) {
		t.Error("expected UpdatedAt to be after CreatedAt")
	}
}

// @audit-ok // deve criar uma chave de API única
func TestGenerateAPIKey(t *testing.T) {
	// Chaves de API geradas
	apiKey1 := generateAPIKey()
	apiKey2 := generateAPIKey()

	// Verifica se as chaves de API são únicas
	if apiKey1 == apiKey2 {
		t.Error("expected API keys to be unique, but they are the same")
	}

	// Verifica se as chaves de API têm o tamanho correto
	if len(apiKey1) != 32 || len(apiKey2) != 32 {
		t.Error("expected API keys to be 32 characters long")
	}
}

// @audit-ok // deve garantir que o mutex está funcionando corretamente ao adicionar saldo
func TestAddBalanceConcurrency(t *testing.T) {

	account := NewAccount("Concurrent User", "concurrent.user@example.com")
	amount := 50.0
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			account.AddBalance(amount)
		}()
	}

	wg.Wait()

	expectedBalance := amount * 10
	if account.Balance != expectedBalance {
		t.Errorf("expected Balance to be %f, got %f", expectedBalance, account.Balance)
	}
}
