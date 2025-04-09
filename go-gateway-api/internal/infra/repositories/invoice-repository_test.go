package repositories_test

import (
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/domain/domainEntities"
	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/infra/repositories"
	"github.com/stretchr/testify/assert"
)

func TestCreateInvoice(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockInvoice := &domainEntities.InvoiceDomain{
		ID:             "123",
		AccountID:      "acc123",
		Amount:         100.0,
		Status:         "pending",
		Description:    "Test invoice",
		CardLastDigits: "1234",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO invoices")).
		ExpectExec().
		WithArgs(mockInvoice.ID, mockInvoice.AccountID, mockInvoice.Amount, mockInvoice.Status, mockInvoice.Description, mockInvoice.CardLastDigits, mockInvoice.CreatedAt, mockInvoice.UpdatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := repositories.NewInvoiceRepository(db)
	err = repo.CreateInvoice(mockInvoice)
	assert.NoError(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestFindByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	createdAt := time.Now()
	updatedAt := time.Now()

	rows := sqlmock.NewRows([]string{"id", "account_id", "amount", "status", "description", "created_at", "updated_at"}).
		AddRow("123", "acc123", 100.0, "pending", "Test invoice", createdAt, updatedAt)

	mock.ExpectPrepare(regexp.QuoteMeta("SELECT id, account_id, amount, status, description, created_at, updated_at FROM invoices WHERE id = $1")).
		ExpectQuery().
		WithArgs("123").
		WillReturnRows(rows)

	repo := repositories.NewInvoiceRepository(db)
	invoice, err := repo.FindByID("123")

	assert.NoError(t, err)
	assert.NotNil(t, invoice)
	assert.Equal(t, "123", invoice.ID)
	assert.Equal(t, "acc123", invoice.AccountID)
	assert.Equal(t, 100.0, invoice.Amount)
	assert.Equal(t, domainEntities.Status("pending"), invoice.Status)
	assert.Equal(t, "Test invoice", invoice.Description)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestFindByAccountID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	createdAt := time.Now()
	updatedAt := time.Now()

	rows := sqlmock.NewRows([]string{"id", "account_id", "amount", "status", "description", "created_at", "updated_at"}).
		AddRow("123", "acc123", 100.0, "pending", "Invoice 1", createdAt, updatedAt).
		AddRow("456", "acc123", 200.0, "paid", "Invoice 2", createdAt, updatedAt)

	mock.ExpectPrepare(regexp.QuoteMeta("SELECT id, account_id, amount, status, description, created_at, updated_at FROM invoices WHERE account_id = $1")).
		ExpectQuery().
		WithArgs("acc123").
		WillReturnRows(rows)

	repo := repositories.NewInvoiceRepository(db)
	invoices, err := repo.FindByAccountID("acc123")

	assert.NoError(t, err)
	assert.Len(t, invoices, 2)
	assert.Equal(t, "123", invoices[0].ID)
	assert.Equal(t, "456", invoices[1].ID)
	assert.Equal(t, domainEntities.Status("pending"), invoices[0].Status)
	assert.Equal(t, domainEntities.Status("paid"), invoices[1].Status)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestUpdateStatus(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockInvoice := &domainEntities.InvoiceDomain{
		ID:        "123",
		Status:    "paid",
		UpdatedAt: time.Now(),
	}

	mock.ExpectPrepare(regexp.QuoteMeta("UPDATE invoices SET status = $1, updated_at = $2 WHERE id = $3")).
		ExpectExec().
		WithArgs(mockInvoice.Status, mockInvoice.UpdatedAt, mockInvoice.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repo := repositories.NewInvoiceRepository(db)
	err = repo.UpdateStatus(mockInvoice)

	assert.NoError(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestCreateInvoice_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mockInvoice := &domainEntities.InvoiceDomain{
		ID:          "123",
		AccountID:   "acc123",
		Amount:      100.0,
		Status:      "pending",
		Description: "Test invoice",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO invoices")).
		ExpectExec().
		WithArgs(mockInvoice.ID, mockInvoice.AccountID, mockInvoice.Amount, mockInvoice.Status, mockInvoice.Description, mockInvoice.CreatedAt, mockInvoice.UpdatedAt).
		WillReturnError(sql.ErrConnDone)

	repo := repositories.NewInvoiceRepository(db)
	db.Close()
	err = repo.CreateInvoice(mockInvoice)
	assert.Error(t, err)
	assert.Equal(t, sql.ErrConnDone, err)
	assert.NotNil(t, mock.ExpectationsWereMet())

}

func TestFindByID_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectPrepare(regexp.QuoteMeta("SELECT id, account_id, amount, status, description, created_at, updated_at FROM invoices WHERE id = $1")).
		ExpectQuery().
		WithArgs("nonexistent").
		WillReturnError(sql.ErrNoRows)

	repo := repositories.NewInvoiceRepository(db)
	invoice, err := repo.FindByID("nonexistent")

	assert.Error(t, err)
	assert.Equal(t, sql.ErrNoRows, err)
	assert.Nil(t, invoice)
	assert.Nil(t, mock.ExpectationsWereMet())
}
