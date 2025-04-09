package service_test

import (
	"fmt"
	"testing"

	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/application/dto"
	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/application/service"
	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/infra/mockRepositories"
)

func TestInvoiceService(t *testing.T) {
	t.Run("should create an invoice successfully", func(t *testing.T) {
		// Arrange
		repo := mockRepositories.NewInMemoryInvoiceRepository()
		accRepo := mockRepositories.NewInMemoryAccountRepository()
		accService := service.NewAccountService(accRepo)
		invoiceService := service.NewInvoiceService(repo, accService)

		// Create account first
		accountInput := &dto.CreateAccountInputDTO{
			Name:  "John Doe",
			Email: "john@example.com",
		}
		account, err := accService.CreateAccount(accountInput)
		if err != nil {
			t.Fatalf("failed to create test account: %v", err)
		}

		input := dto.CreateInvoiceInputDTO{
			APIKey:          account.APIKey,
			Amount:          100.0,
			Description:     "Test invoice",
			PaymentType:     "credit_card",
			CardNumber:      "4111111111111111",
			CVV:             "123",
			ExpirationMonth: 12,
			ExpirationYear:  2030,
			CardHolderName:  "John Doe",
		}

		// Act
		invoice, err := invoiceService.CreateInvoice(input, account.APIKey)
		fmt.Println("ITS ALAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIVE !:", invoice)
		// Assert
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if invoice == nil {
			t.Fatal("expected invoice, got nil")
		}
		if invoice.Amount != input.Amount {
			t.Errorf("expected amount %f, got %f", input.Amount, invoice.Amount)
		}
		if invoice.Description != input.Description {
			t.Errorf("expected description %s, got %s", input.Description, invoice.Description)
		}
		if invoice.Status != "pending" {
			t.Errorf("expected status 'pending', got %s", invoice.Status)
		}
		if invoice.AccountID != account.ID {
			t.Errorf("expected account ID %s, got %s", account.ID, invoice.AccountID)
		}
	})

	t.Run("should fail to create invoice with invalid account", func(t *testing.T) {
		// Arrange
		repo := mockRepositories.NewInMemoryInvoiceRepository()
		accRepo := mockRepositories.NewInMemoryAccountRepository()
		accService := service.NewAccountService(accRepo)
		invoiceService := service.NewInvoiceService(repo, accService)

		input := dto.CreateInvoiceInputDTO{
			APIKey:          "invalid-api-key",
			Amount:          100.0,
			Description:     "Test invoice",
			PaymentType:     "credit_card",
			CardNumber:      "4111111111111111",
			CVV:             "123",
			ExpirationMonth: 12,
			ExpirationYear:  2030,
			CardHolderName:  "John Doe",
		}

		// Act
		_, err := invoiceService.CreateInvoice(input, input.APIKey)

		// Assert
		if err == nil {
			t.Errorf("expected error for non-existent account, got nil")
		}
	})

	t.Run("should find invoice by ID", func(t *testing.T) {
		// Arrange
		repo := mockRepositories.NewInMemoryInvoiceRepository()
		accRepo := mockRepositories.NewInMemoryAccountRepository()
		accService := service.NewAccountService(accRepo)
		invoiceService := service.NewInvoiceService(repo, accService)
		accountInput := &dto.CreateAccountInputDTO{
			Name:  "Alice",
			Email: "alice@example.com",
		}
		account, _ := accService.CreateAccount(accountInput)

		// Create invoice
		input := dto.CreateInvoiceInputDTO{
			APIKey:          account.APIKey,
			Amount:          150.0,
			Description:     "Test invoice",
			PaymentType:     "credit_card",
			CardNumber:      "4111111111111111",
			CVV:             "123",
			ExpirationMonth: 12,
			ExpirationYear:  2030,
			CardHolderName:  "Alice",
		}
		created, err := invoiceService.CreateInvoice(input, account.APIKey)
		if err != nil {
			t.Fatalf("failed to create test invoice: %v", err)
		}

		// Act
		found, err := invoiceService.FindInvoiceByID(created.ID, account.APIKey)

		// Assert
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if found.ID != created.ID {
			t.Errorf("expected ID %s, got %s", created.ID, found.ID)
		}
		if found.Amount != created.Amount {
			t.Errorf("expected amount %f, got %f", created.Amount, found.Amount)
		}
	})

	t.Run("should not find non-existent invoice", func(t *testing.T) {
		// Arrange
		repo := mockRepositories.NewInMemoryInvoiceRepository()
		accRepo := mockRepositories.NewInMemoryAccountRepository()
		accService := service.NewAccountService(accRepo)
		invoiceService := service.NewInvoiceService(repo, accService)

		// Act
		_, err := invoiceService.FindInvoiceByID("non-existent-id", "valid-api-key")

		// Assert
		if err == nil {
			t.Errorf("expected error for non-existent invoice ID")
		}
	})

	t.Run("should list all invoices for an account", func(t *testing.T) {
		// Arrange
		repo := mockRepositories.NewInMemoryInvoiceRepository()
		accRepo := mockRepositories.NewInMemoryAccountRepository()
		accService := service.NewAccountService(accRepo)
		invoiceService := service.NewInvoiceService(repo, accService)
		accountInput := &dto.CreateAccountInputDTO{
			Name:  "Bob",
			Email: "bob@example.com",
		}
		account, _ := accService.CreateAccount(accountInput)

		// Create multiple invoices
		for i := 0; i < 3; i++ {
			input := dto.CreateInvoiceInputDTO{
				APIKey:          account.APIKey,
				Amount:          100.0 * float64(i+1),
				Description:     "Test invoice",
				PaymentType:     "credit_card",
				CardNumber:      "4111111111111111",
				CVV:             "123",
				ExpirationMonth: 12,
				ExpirationYear:  2030,
				CardHolderName:  "Bob",
			}
			_, _ = invoiceService.CreateInvoice(input, account.APIKey)
		}

		// Act
		invoices, err := invoiceService.ListInvoicesByAccount(account.ID)

		// Assert
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if len(invoices) != 3 {
			t.Errorf("expected 3 invoices, got %d", len(invoices))
		}
	})

}
