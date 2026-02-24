package repository

import (
	"go-web/migrations"

	"gorm.io/gorm"
)

type AccountRepository interface {
	Create(contact *migrations.Account) error
	Update(contact *migrations.Account) error
	Delete(contact *migrations.Account) error
	Select(contact *migrations.Account) error
	SelectAll() ([]migrations.Account, error)
}

type accountRepo struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepo{db: db}
}

func (r *accountRepo) Create(contact *migrations.Account) error {
	return r.db.Create(contact).Error
}

func (r *accountRepo) Update(contact *migrations.Account) error {
	return r.db.Save(contact).Error
}

func (r *accountRepo) SelectAll() ([]migrations.Account, error) {
	var accounts []migrations.Account
	err := r.db.Find(&accounts).Error
	return accounts, err
}

func (r *accountRepo) Delete(contact *migrations.Account) error {
	return r.db.Delete(contact).Error
}

func (r *accountRepo) Select(contact *migrations.Account) error {
	return r.db.First(contact).Error
}
