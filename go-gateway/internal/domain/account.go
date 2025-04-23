package domain

import (
	"crypto/rand"
	"encoding/hex"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        string
	Name      string
	Email     string
	API_KEY   string
	Balance   float64
	Mu        sync.RWMutex
	CreatedAt time.Time
	UpdatedAt time.Time
}

func createAPIKey() string {
	arr := make([]byte, 16)
	rand.Read(arr)
	//
	return hex.EncodeToString(arr)
}

func CreateAccount(name, email string) *Account {
	account := &Account{
		Name:      name,
		Email:     email,
		ID:        uuid.New().String(),
		Balance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		API_KEY:   createAPIKey(),
	}

	return account
}

func (act *Account) SetBalance(amount float64) {
	// before change value, freeze the values for not create a racing condition
	act.Mu.Lock()
	// unlock freeze to the account object;
	// defer will await to the end of the function to execute
	defer act.Mu.Unlock()
	// now change account balance
	act.Balance += amount
	act.UpdatedAt = time.Now()
}
