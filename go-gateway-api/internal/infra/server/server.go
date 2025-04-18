package server

import (
	"net/http"

	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/application/service"
	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/infra/handlers"
	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/infra/middleware"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	router         *chi.Mux
	server         *http.Server
	accountService *service.AccountService
	invoiceService *service.InvoiceService
	port           string
}

func NewServer(accountService *service.AccountService, invoiceService *service.InvoiceService, port string) *Server {
	router := chi.NewRouter()
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	return &Server{
		router:         router,
		server:         srv,
		accountService: accountService,
		invoiceService: invoiceService,
		port:           port,
	}
}

func (s *Server) SetupRoutes() {
	accountHandler := handlers.NewAccountHandler(s.accountService)
	invoiceHandler := handlers.NewInvoiceHandler(s.invoiceService)
	authMiddleware := middleware.NewAuthMiddleware(s.accountService)

	s.router.Post("/accounts", accountHandler.Create)
	s.router.Get("/accounts", accountHandler.Get)

	s.router.Group(func(r chi.Router) {
		r.Use(authMiddleware.Authenticate)
		s.router.Post("/invoices", invoiceHandler.Create)
		s.router.Get("/invoices", invoiceHandler.GetByAccountApiKey)
		s.router.Get("/invoices/{id}", invoiceHandler.GetByID)
	})
}

func (s *Server) Start() error {
	s.server = &http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,
	}

	return s.server.ListenAndServe()
}
