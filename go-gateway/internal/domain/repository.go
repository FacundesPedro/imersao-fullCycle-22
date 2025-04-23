package domain

type AccountRepository interface {
	Save(account *Account) error
	FindByAPIKey(key string) (*Account, error)
	FindByID(id string) (*Account, error)
	SetBalance(account *Account) error
}

type InvoiceRepository interface {
	Save(invoice *Invoice) error
	FindByID(id string) (*Invoice, error)
	FindByAccountID(accountID string) ([]*Invoice, error)
	UpdateStatus(invoice *Invoice) error
}
