package repository

import (
	"database/sql"
	"time"

	"github.com/devfullcycle/imersao22/go-gateway/internal/domain"
)

// AccountRepository implementa operações de persistência para Account
type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (repo *AccountRepository) Save(account *domain.Account) error {
	stmt, err := repo.db.Prepare(`
		INSERT INTO accounts (id, name, email, api_key, balance, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`)
	if err != nil {
		return err
	}
	//
	defer stmt.Close()
	//
	_, err = stmt.Exec(
		account.ID,
		account.Name,
		account.Email,
		account.API_KEY,
		account.Balance,
		account.CreatedAt,
		account.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (repo *AccountRepository) FindByAPIKey(key string) (*domain.Account, error) {
	var act domain.Account
	var createdAt, updateAt time.Time

	err := repo.db.QueryRow(`
		SELECT id, name, email, api_key, balance, created_at, updated_at
		FROM accounts
		WHERE api_key = $1	
	`, key).Scan(
		&act.ID,
		&act.Name,
		&act.Email,
		&act.API_KEY,
		&act.Balance,
		&createdAt,
		&updateAt,
	)
	if err == sql.ErrNoRows {
		return nil, domain.ErrAccountNotFound
	}
	act.CreatedAt = createdAt
	act.UpdatedAt = updateAt

	return &act, nil
}

func (repo *AccountRepository) SetBalance(act *domain.Account) error {
	var currentBalance float64
	//
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	err = tx.QueryRow(`
		SELECT balance FROM accounts WHERE id = $1 FOR UPDATE
	`, act.ID).Scan(&currentBalance)
	if err == sql.ErrNoRows {
		return domain.ErrAccountNotFound
	}
	if err != nil {
		return err
	}
	_, err = tx.Exec(`
		UPDATE accounts
		SET balance = $1, updated_at = $2
		WHERE id = $3
	`, act.Balance, time.Now(), act.ID)
	if err != nil {
		return err
	}

	return tx.Commit()
}
func (r *AccountRepository) FindByID(id string) (*domain.Account, error) {
	var account domain.Account
	var createdAt, updatedAt time.Time

	err := r.db.QueryRow(`
		SELECT id, name, email, api_key, balance, created_at, updated_at
		FROM accounts
		WHERE id = $1
	`, id).Scan(
		&account.ID,
		&account.Name,
		&account.Email,
		&account.API_KEY,
		&account.Balance,
		&createdAt,
		&updatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, domain.ErrAccountNotFound
	}
	if err != nil {
		return nil, err
	}

	account.CreatedAt = createdAt
	account.UpdatedAt = updatedAt
	return &account, nil
}
