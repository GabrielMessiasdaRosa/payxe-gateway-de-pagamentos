package domain

import (
	"crypto/rand"
	"encoding/hex"
	"sync"
	"time"

	"github.com/google/uuid"
)


type Account struct {
	ID 	 string
	Name string
	Email string
	APIKey string
	Balance float64
	mutex sync.Mutex
	CreatedAt time.Time
	UpdatedAt time.Time
}

func generateAPIKey() string {
	// b é um slice de bytes com 16 bytes
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}


// * siginifica um ponteiro
func NewAccount(name, email string) *Account {
	account := &Account{
		ID: uuid.New().String(),
		Name: name,
		Email: email,
		APIKey: generateAPIKey(),
		Balance: 0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return account
}

func (account *Account) AddBalance(amount float64) {
	account.mutex.Lock()
	defer account.mutex.Unlock()
	account.Balance += amount
	account.UpdatedAt = time.Now()
}

