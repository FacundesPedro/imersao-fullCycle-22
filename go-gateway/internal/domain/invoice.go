package domain

import (
	"time"

	"github.com/devfullcycle/imersao22/go-gateway/internal/domain"
)

// Status contants
type Status string

const (
	StatusPending  Status = "pending"
	StatusApproved Status = "approved"
	StatusRejected Status = "rejected"
)

type Invoice struct {
	ID             string
	AccountID      string
	Amount         float64
	Status         Status
	Description    string
	PaymentType    string
	CardLastDigits string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type CreditCard struct {
	Number      string
	Cvv         string
	ExpiryMonth int
	ExpiryYear  int
	HolderName  string
}

func NewInvoice(accountId string, amount float64, description string, paymentType string, card *CreditCard) (*Invoice, error) {
	if amount <= 0 {
		return nil, domain.ErrInvalidAmount
	}

	return nil
}
