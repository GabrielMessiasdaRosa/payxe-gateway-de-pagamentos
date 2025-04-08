package domainEntities

import (
	"sync"
	"time"

	"github.com/google/uuid"
)

type AccountDomain struct {
	ID        string
	Name      string
	Email     string
	APIKey    string
	Balance   float64
	mutex     sync.Mutex
	CreatedAt time.Time
	UpdatedAt time.Time
}

// * siginifica um ponteiro
func NewAccount(name, email string) *AccountDomain {
	account := &AccountDomain{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		APIKey:    uuid.New().String(),
		Balance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return account
}

func (account *AccountDomain) AddBalance(amount float64) {
	account.mutex.Lock()
	defer account.mutex.Unlock()
	account.Balance += amount
	account.UpdatedAt = time.Now()
}
