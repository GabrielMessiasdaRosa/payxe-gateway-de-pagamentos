package main

import (
	"log"

	// importe o lib

	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/databaseConn"
	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/application/service"
	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/infra/repositories"
	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/infra/server"
	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/utils"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	pgDb, err := databaseConn.ConnectToPostgres()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer pgDb.Close()

	accountRepository := repositories.NewAccountRepository(pgDb)
	accountService := service.NewAccountService(accountRepository)

	port := utils.GetEnv("PORT", "8080")
	log.Printf("Server is running on port: %s", port)
	srv := server.NewServer(accountService, port)
	srv.SetupRoutes()
	if err := srv.Start(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	log.Println("Server stopped")

}
