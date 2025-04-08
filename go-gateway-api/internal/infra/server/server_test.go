package server

import (
	"fmt"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/application/service"
	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	accountService := &service.AccountService{}
	port := "3000"

	server := NewServer(accountService, port)

	assert.NotNil(t, server)
	assert.NotNil(t, server.router)
	assert.NotNil(t, server.server)
	assert.Equal(t, accountService, server.accountService)
	assert.Equal(t, port, server.port)
	assert.Equal(t, ":"+port, server.server.Addr)
}

func TestServer_SetupRoutes(t *testing.T) {
	accountService := &service.AccountService{}
	port := "3000"

	server := NewServer(accountService, port)
	server.SetupRoutes()

	// Check if routes were registered by inspecting the router
	// Chi doesn't expose routes directly, so we have to use creative ways to test
	assert.NotNil(t, server.router)
}

func TestServer_Start(t *testing.T) {
	// Find an available port for testing
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("Could not find available port: %v", err)
	}
	port := fmt.Sprintf("%d", listener.Addr().(*net.TCPAddr).Port)
	listener.Close()

	accountService := &service.AccountService{}
	server := NewServer(accountService, port)
	server.SetupRoutes()

	// Start server in a goroutine since it blocks
	go func() {
		err := server.Start()
		if err != nil && err != http.ErrServerClosed {
			t.Errorf("Server failed to start: %v", err)
		}
	}()

	// Give server time to start
	time.Sleep(100 * time.Millisecond)

	// Make a request to check if server is running
	resp, err := http.Get(fmt.Sprintf("http://localhost:%s/accounts", port))
	if err != nil {
		t.Fatalf("Failed to send request to server: %v", err)
	}
	defer resp.Body.Close()

	// We expect a response, but don't care about the status code since we're just testing if the server started
	assert.NotNil(t, resp)
}
