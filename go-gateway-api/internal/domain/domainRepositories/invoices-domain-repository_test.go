package domainRepositories

import (
	"errors"
	"testing"

	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/domain/domainEntities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockInvoiceDomainRepository is a mock implementation of InvoiceDomainRepository
type MockInvoiceDomainRepository struct {
	mock.Mock
}

func (m *MockInvoiceDomainRepository) CreateInvoice(invoice *domainEntities.InvoiceDomain) error {
	args := m.Called(invoice)
	return args.Error(0)
}

func (m *MockInvoiceDomainRepository) FindByID(id string) (*domainEntities.InvoiceDomain, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domainEntities.InvoiceDomain), args.Error(1)
}

func (m *MockInvoiceDomainRepository) FindByReference(reference string) (*domainEntities.InvoiceDomain, error) {
	args := m.Called(reference)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domainEntities.InvoiceDomain), args.Error(1)
}

func (m *MockInvoiceDomainRepository) UpdateStatus(invoice *domainEntities.InvoiceDomain) error {
	args := m.Called(invoice)
	return args.Error(0)
}

func TestMockInvoiceDomainRepository_CreateInvoice(t *testing.T) {
	mockRepo := new(MockInvoiceDomainRepository)
	invoice := &domainEntities.InvoiceDomain{ID: "123"}

	// Test success case
	mockRepo.On("CreateInvoice", invoice).Return(nil)
	err := mockRepo.CreateInvoice(invoice)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)

	// Test error case
	expectedErr := errors.New("database error")
	mockRepo = new(MockInvoiceDomainRepository)
	mockRepo.On("CreateInvoice", invoice).Return(expectedErr)
	err = mockRepo.CreateInvoice(invoice)
	assert.Equal(t, expectedErr, err)
	mockRepo.AssertExpectations(t)
}

func TestMockInvoiceDomainRepository_FindByID(t *testing.T) {
	mockRepo := new(MockInvoiceDomainRepository)
	invoice := &domainEntities.InvoiceDomain{ID: "123"}

	// Test success case
	mockRepo.On("FindByID", "123").Return(invoice, nil)
	result, err := mockRepo.FindByID("123")
	assert.NoError(t, err)
	assert.Equal(t, invoice, result)
	mockRepo.AssertExpectations(t)

	// Test not found case
	mockRepo = new(MockInvoiceDomainRepository)
	mockRepo.On("FindByID", "456").Return(nil, errors.New("invoice not found"))
	result, err = mockRepo.FindByID("456")
	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestMockInvoiceDomainRepository_UpdateStatus(t *testing.T) {
	mockRepo := new(MockInvoiceDomainRepository)
	invoice := &domainEntities.InvoiceDomain{ID: "123", Status: "PAID"}

	// Test success case
	mockRepo.On("UpdateStatus", invoice).Return(nil)
	err := mockRepo.UpdateStatus(invoice)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)

	// Test error case
	expectedErr := errors.New("database error")
	mockRepo = new(MockInvoiceDomainRepository)
	mockRepo.On("UpdateStatus", invoice).Return(expectedErr)
	err = mockRepo.UpdateStatus(invoice)
	assert.Equal(t, expectedErr, err)
	mockRepo.AssertExpectations(t)
}
