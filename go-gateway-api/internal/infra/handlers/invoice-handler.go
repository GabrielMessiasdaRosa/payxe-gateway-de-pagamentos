// filepath: /home/gmrosa/Desktop/Estudo/imersoes/payxe-gateway-de-pagamentos/go-gateway-api/internal/infra/handlers/invoice-handler.go
package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/application/dto"
	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/application/service"
	"github.com/go-chi/chi/v5"
)

type InvoiceHandler struct {
	InvoiceService *service.InvoiceService
}

func NewInvoiceHandler(invoiceService *service.InvoiceService) *InvoiceHandler {
	return &InvoiceHandler{
		InvoiceService: invoiceService,
	}
}

func (h *InvoiceHandler) Create(w http.ResponseWriter, r *http.Request) {
	input := dto.CreateInvoiceInputDTO{}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if input.Amount <= 0 || input.Description == "" || input.AccountID == "" {
		http.Error(w, "Missing or invalid fields", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	apiKey := r.Header.Get("X-API-Key")
	invoice, err := h.InvoiceService.CreateInvoice(input, apiKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(invoice)
}

func (h *InvoiceHandler) GetByAccountApiKey(w http.ResponseWriter, r *http.Request) {
	if h.InvoiceService == nil {
		http.Error(w, "Invoice service is not initialized", http.StatusInternalServerError)
		return
	}

	apiKey := r.Header.Get("X-API-Key")
	if apiKey == "" {
		http.Error(w, "API key is required", http.StatusBadRequest)
		return
	}

	account, err := h.InvoiceService.AccountService.FindByAPIKey(apiKey)
	if err != nil {
		http.Error(w, "Invalid API key", http.StatusUnauthorized)
		return
	}

	invoices, err := h.InvoiceService.ListInvoicesByAccount(account.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(invoices)
}

func (h *InvoiceHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	if h.InvoiceService == nil {
		http.Error(w, "Invoice service is not initialized", http.StatusInternalServerError)
		return
	}

	invoiceID := chi.URLParam(r, "id")
	if invoiceID == "" {
		http.Error(w, "Invoice ID is required", http.StatusBadRequest)
		return
	}

	apiKey := r.Header.Get("X-API-Key")
	if apiKey == "" {
		http.Error(w, "API key is required", http.StatusBadRequest)
		return
	}

	invoice, err := h.InvoiceService.FindInvoiceByID(invoiceID, apiKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(invoice)
}
