package repositories

import (
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/domain/domainEntities"
	"github.com/stretchr/testify/assert"
)

// @audit-ok // deve salvar uma conta no banco de dados
func TestAccountRepository_Save(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewAccountRepository(db)

	account := &domainEntities.AccountDomain{
		ID:        "1",
		Name:      "Test User",
		Email:     "test@example.com",
		APIKey:    "test-api-key",
		Balance:   100.0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mock.ExpectPrepare("INSERT INTO accounts").
		ExpectExec().
		WithArgs(account.ID, account.Name, account.Email, account.APIKey, account.Balance, account.CreatedAt, account.UpdatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.CreateAccount(account)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

// @audit-ok // deve encontrar uma conta pelo ID
func TestAccountRepository_FindByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewAccountRepository(db)

	account := &domainEntities.AccountDomain{
		ID:        "1",
		Name:      "Test User",
		Email:     "test@example.com",
		APIKey:    "test-api-key",
		Balance:   100.0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "api_key", "balance", "created_at", "updated_at"}).
		AddRow(account.ID, account.Name, account.Email, account.APIKey, account.Balance, account.CreatedAt, account.UpdatedAt)

	mock.ExpectPrepare("SELECT id, name, email, api_key, balance, created_at, updated_at FROM accounts WHERE id = \\$1").
		ExpectQuery().
		WithArgs(account.ID).
		WillReturnRows(rows)

	result, err := repo.FindByID(account.ID)
	assert.NoError(t, err)
	assert.Equal(t, account, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

// @audit-ok // deve encontrar uma conta pelo API Key
func TestAccountRepository_FindByAPIKey(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewAccountRepository(db)

	account := &domainEntities.AccountDomain{
		ID:        "1",
		Name:      "Test User",
		Email:     "test@example.com",
		APIKey:    "test-api-key",
		Balance:   100.0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "api_key", "balance", "created_at", "updated_at"}).
		AddRow(account.ID, account.Name, account.Email, account.APIKey, account.Balance, account.CreatedAt, account.UpdatedAt)

	mock.ExpectPrepare("SELECT id, name, email, api_key, balance, created_at, updated_at FROM accounts WHERE api_key = \\$1").
		ExpectQuery().
		WithArgs(account.APIKey).
		WillReturnRows(rows)

	result, err := repo.FindByAPIKey(account.APIKey)
	assert.NoError(t, err)
	assert.Equal(t, account, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

// @audit-ok // deve atualizar o saldo da conta
func TestAccountRepository_UpdateBalance(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewAccountRepository(db)

	account := &domainEntities.AccountDomain{
		ID:        "1",
		Balance:   50.0,
		UpdatedAt: time.Now(),
	}

	currentBalance := 100.0

	mock.ExpectBegin()

	mock.ExpectQuery("SELECT balance FROM accounts WHERE id = \\$1 FOR UPDATE").
		WithArgs(account.ID).
		WillReturnRows(sqlmock.NewRows([]string{"balance"}).AddRow(currentBalance))

	mock.ExpectExec("UPDATE accounts SET balance = \\$1, updated_at = \\$2 WHERE id = \\$3").
		WithArgs(currentBalance+account.Balance, account.UpdatedAt, account.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	err = repo.UpdateBalance(account)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

// @audit-ok // deve retornar erro se não encontrar a conta
func TestAccountRepository_UpdateBalance_NoRows(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewAccountRepository(db)

	account := &domainEntities.AccountDomain{
		ID:        "1",
		Balance:   50.0,
		UpdatedAt: time.Now(),
	}

	mock.ExpectBegin()

	mock.ExpectQuery("SELECT balance FROM accounts WHERE id = \\$1 FOR UPDATE").
		WithArgs(account.ID).
		WillReturnError(sql.ErrNoRows)

	mock.ExpectRollback()

	err = repo.UpdateBalance(account)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

// @audit-ok // deve retornar erro se houver erro na consulta
func TestAccountRepository_UpdateBalance_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewAccountRepository(db)

	account := &domainEntities.AccountDomain{
		ID:        "1",
		Balance:   50.0,
		UpdatedAt: time.Now(),
	}

	mock.ExpectBegin()

	mock.ExpectQuery("SELECT balance FROM accounts WHERE id = \\$1 FOR UPDATE").
		WithArgs(account.ID).
		WillReturnError(sql.ErrConnDone)

	mock.ExpectRollback()

	err = repo.UpdateBalance(account)
	assert.Error(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

// @audit-ok // deve retornar erro se houver erro na execução
func TestAccountRepository_UpdateBalance_ExecError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewAccountRepository(db)

	account := &domainEntities.AccountDomain{
		ID:        "1",
		Balance:   50.0,
		UpdatedAt: time.Now(),
	}

	currentBalance := 100.0

	mock.ExpectBegin()

	mock.ExpectQuery("SELECT balance FROM accounts WHERE id = \\$1 FOR UPDATE").
		WithArgs(account.ID).
		WillReturnRows(sqlmock.NewRows([]string{"balance"}).AddRow(currentBalance))

	mock.ExpectExec("UPDATE accounts SET balance = \\$1, updated_at = \\$2 WHERE id = \\$3").
		WithArgs(currentBalance+account.Balance, account.UpdatedAt, account.ID).
		WillReturnError(sql.ErrConnDone)

	mock.ExpectRollback()

	err = repo.UpdateBalance(account)
	assert.Error(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

// @audit-ok // deve retornar erro se houver erro no commit
func TestAccountRepository_UpdateBalance_CommitError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewAccountRepository(db)

	account := &domainEntities.AccountDomain{
		ID:        "1",
		Balance:   50.0,
		UpdatedAt: time.Now(),
	}

	currentBalance := 100.0

	mock.ExpectBegin()

	mock.ExpectQuery("SELECT balance FROM accounts WHERE id = \\$1 FOR UPDATE").
		WithArgs(account.ID).
		WillReturnRows(sqlmock.NewRows([]string{"balance"}).AddRow(currentBalance))

	mock.ExpectExec("UPDATE accounts SET balance = \\$1, updated_at = \\$2 WHERE id = \\$3").
		WithArgs(currentBalance+account.Balance, account.UpdatedAt, account.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit().WillReturnError(sql.ErrTxDone)

	err = repo.UpdateBalance(account)
	assert.Error(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
