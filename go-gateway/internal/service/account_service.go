package service

import (
	"github.com/facundes/imersao22/go-gateway/internal/domain"
	"github.com/facundes/imersao22/go-gateway/internal/dto"
)

type AccountService struct {
	repository domain.AccountRepository
}

func NewAccountService(repository domain.AccountRepository) *AccountService {
	return &AccountService{repository: repository}
}

func (s *AccountService) NewAccount(input dto.CreateAccountInput) (*dto.AccountOutput, error) {
	act := dto.ToAccount(input)

	prevAccount, err := s.repository.FindByAPIKey(act.API_KEY)
	if err != nil && err != domain.ErrAccountNotFound {
		return nil, err
	}

	if prevAccount != nil {
		return nil, domain.ErrDuplicatedAPIKey
	}
	err = s.repository.Save(act)
	if err != nil {
		return nil, err
	}

	out := dto.FromAccount(act)
	return &out, nil
}

func (s *AccountService) SetBalance(apiKey string, amount float64) (*dto.AccountOutput, error) {
	act, err := s.repository.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	act.SetBalance(amount)
	err = s.repository.SetBalance(act)
	if err != nil {
		return nil, err
	}

	out := dto.FromAccount(act)
	return &out, nil
}

func (s *AccountService) FindByAPIKey(apiKey string) (*dto.AccountOutput, error) {
	act, err := s.repository.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	out := dto.FromAccount(act)
	return &out, nil
}
