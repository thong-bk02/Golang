package service

import (
	"go-web/internal/repository"
	"go-web/migrations"
)

type AccountService interface {
	SubmitForm(input migrations.Account) error
	GetAllAccounts() ([]migrations.Account, error)
	Delete(id uint) error
	Select(id uint) (*migrations.Account, error)
}

type accountService struct {
	repo repository.AccountRepository
}

func NewAccountService(repo repository.AccountRepository) AccountService {
	return &accountService{repo: repo}
}

func (s *accountService) SubmitForm(input migrations.Account) error {
	// Bạn có thể thêm logic kiểm tra, gửi email thông báo... ở đây
	return s.repo.Create(&input)
}

func (s *accountService) GetAllAccounts() ([]migrations.Account, error) {
	return s.repo.SelectAll()
}

func (s *accountService) Delete(id uint) error {
	account := migrations.Account{ID: id}
	return s.repo.Delete(&account)
}

func (s *accountService) Select(id uint) (*migrations.Account, error) {
	account := &migrations.Account{ID: id}
	err := s.repo.Select(account)
	return account, err
}
