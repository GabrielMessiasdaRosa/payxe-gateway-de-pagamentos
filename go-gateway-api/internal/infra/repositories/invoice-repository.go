package repositories

import (
	"database/sql"

	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/domain/domainEntities"
)

type InvoiceRepository struct {
	db *sql.DB
}

func NewInvoiceRepository(db *sql.DB) *InvoiceRepository {
	return &InvoiceRepository{
		db: db,
	}
}

func (r *InvoiceRepository) CreateInvoice(invoice *domainEntities.InvoiceDomain) error {
	stmt, err := r.db.Prepare("INSERT INTO invoices (id, account_id, amount, status, description, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(invoice.ID, invoice.AccountID, invoice.Amount, invoice.Status, invoice.Description, invoice.CreatedAt, invoice.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *InvoiceRepository) FindByID(id string) (*domainEntities.InvoiceDomain, error) {
	stmt, err := r.db.Prepare("SELECT id, account_id, amount, status, description, created_at, updated_at FROM invoices WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	invoice := &domainEntities.InvoiceDomain{}
	err = stmt.QueryRow(id).Scan(&invoice.ID, &invoice.AccountID, &invoice.Amount, &invoice.Status, &invoice.Description, &invoice.CreatedAt, &invoice.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return invoice, nil
}

func (r *InvoiceRepository) FindByAccountID(accountID string) ([]*domainEntities.InvoiceDomain, error) {
	stmt, err := r.db.Prepare("SELECT id, account_id, amount, status, description, created_at, updated_at FROM invoices WHERE account_id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	invoices := []*domainEntities.InvoiceDomain{}
	for rows.Next() {
		invoice := &domainEntities.InvoiceDomain{}
		err = rows.Scan(&invoice.ID, &invoice.AccountID, &invoice.Amount, &invoice.Status, &invoice.Description, &invoice.CreatedAt, &invoice.UpdatedAt)
		if err != nil {
			return nil, err
		}
		invoices = append(invoices, invoice)
	}
	return invoices, nil
}

func (r *InvoiceRepository) UpdateStatus(invoice *domainEntities.InvoiceDomain) error {
	stmt, err := r.db.Prepare("UPDATE invoices SET status = $1, updated_at = $2 WHERE id = $3")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(invoice.Status, invoice.UpdatedAt, invoice.ID)
	if err != nil {
		return err
	}
	return nil
}
