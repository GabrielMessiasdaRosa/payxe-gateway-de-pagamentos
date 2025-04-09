package domainRepositories

import (
	"errors"
	"testing"

	"github.com/GabrielMessiasdaRosa/payxe-gateway-de-pagamentos/go-gateway-api/internal/domain/domainEntities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockAccountDomainRepository is a mock of AccountDomainRepository interface
type MockAccountDomainRepository struct {
	mock.Mock
}

func (m *MockAccountDomainRepository) CreateAccount(account *domainEntities.AccountDomain) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *MockAccountDomainRepository) FindByID(id string) (*domainEntities.AccountDomain, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domainEntities.AccountDomain), args.Error(1)
}

func (m *MockAccountDomainRepository) FindByAPIKey(apiKey string) (*domainEntities.AccountDomain, error) {
	args := m.Called(apiKey)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domainEntities.AccountDomain), args.Error(1)
}

func (m *MockAccountDomainRepository) UpdateBalance(account *domainEntities.AccountDomain) error {
	args := m.Called(account)
	return args.Error(0)
}

func TestMockAccountDomainRepository_CreateAccount(t *testing.T) {
	mockRepo := new(MockAccountDomainRepository)
	account := &domainEntities.AccountDomain{ID: "123"}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("CreateAccount", account).Return(nil).Once()
		err := mockRepo.CreateAccount(account)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		expectedErr := errors.New("creation error")
		mockRepo.On("CreateAccount", account).Return(expectedErr).Once()
		err := mockRepo.CreateAccount(account)
		assert.Equal(t, expectedErr, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestMockAccountDomainRepository_FindByID(t *testing.T) {
	mockRepo := new(MockAccountDomainRepository)
	account := &domainEntities.AccountDomain{ID: "123"}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("FindByID", "123").Return(account, nil).Once()
		result, err := mockRepo.FindByID("123")
		assert.Nil(t, err)
		assert.Equal(t, account, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("not found", func(t *testing.T) {
		expectedErr := errors.New("not found")
		mockRepo.On("FindByID", "456").Return(nil, expectedErr).Once()
		result, err := mockRepo.FindByID("456")
		assert.Equal(t, expectedErr, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestMockAccountDomainRepository_FindByAPIKey(t *testing.T) {
	mockRepo := new(MockAccountDomainRepository)
	account := &domainEntities.AccountDomain{ID: "123", APIKey: "api-key-123"}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("FindByAPIKey", "api-key-123").Return(account, nil).Once()
		result, err := mockRepo.FindByAPIKey("api-key-123")
		assert.Nil(t, err)
		assert.Equal(t, account, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("not found", func(t *testing.T) {
		expectedErr := errors.New("not found")
		mockRepo.On("FindByAPIKey", "invalid-key").Return(nil, expectedErr).Once()
		result, err := mockRepo.FindByAPIKey("invalid-key")
		assert.Equal(t, expectedErr, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestMockAccountDomainRepository_UpdateBalance(t *testing.T) {
	mockRepo := new(MockAccountDomainRepository)
	account := &domainEntities.AccountDomain{ID: "123", Balance: 100}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("UpdateBalance", account).Return(nil).Once()
		err := mockRepo.UpdateBalance(account)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		expectedErr := errors.New("update error")
		mockRepo.On("UpdateBalance", account).Return(expectedErr).Once()
		err := mockRepo.UpdateBalance(account)
		assert.Equal(t, expectedErr, err)
		mockRepo.AssertExpectations(t)
	})
}
