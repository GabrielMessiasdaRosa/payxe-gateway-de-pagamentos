package repository

import (
	"database/sql"

	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/domain"
)

type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{
		db: db,
	}
}
func (r *AccountRepository) Save(account *domain.Account) error {
	stmt, err := r.db.Prepare("INSERT INTO accounts (id, name, email, api_key, balance, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	// _ é um identificador de variável que não será utilizado (Blank Identifier)
	_, err = stmt.Exec(account.ID, account.Name, account.Email, account.APIKey, account.Balance, account.CreatedAt, account.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *AccountRepository) FindByID(id string) (*domain.Account, error) {
	stmt, err := r.db.Prepare("SELECT id, name, email, api_key, balance, created_at, updated_at FROM accounts WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	account := &domain.Account{}
	err = stmt.QueryRow(id).Scan(&account.ID, &account.Name, &account.Email, &account.APIKey, &account.Balance, &account.CreatedAt, &account.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (r *AccountRepository) FindByAPIKey(apiKey string) (*domain.Account, error) {
	stmt, err := r.db.Prepare("SELECT id, name, email, api_key, balance, created_at, updated_at FROM accounts WHERE api_key = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	account := &domain.Account{}
	err = stmt.QueryRow(apiKey).Scan(&account.ID, &account.Name, &account.Email, &account.APIKey, &account.Balance, &account.CreatedAt, &account.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (r *AccountRepository) UpdateBalance(account *domain.Account) error {
	
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var currentBalance float64
	err = tx.QueryRow(("SELECT balance FROM accounts WHERE id = ? FOR UPDATE"), account.ID).Scan(&currentBalance)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		return err 
	}

	_, err = tx.Exec("UPDATE accounts SET balance = ?, updated_at = ? WHERE id = ?", currentBalance+account.Balance, account.UpdatedAt, account.ID)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}