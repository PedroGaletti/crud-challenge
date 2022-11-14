package accounts

import (
	"github.com/jinzhu/gorm"
)

const (
	tbAccounts = "accounts"
)

// IAccountRepository: interface of Account repository
type IAccountRepository interface {
	Create(*Account) error
	FindOne(int64) (*Account, error)
}

// AccountRepository: struct of Account repository
type AccountRepository struct {
	db *gorm.DB
}

// NewAccountRepository: create a new Account repository
func NewAccountRepository(db *gorm.DB) IAccountRepository {
	return &AccountRepository{db}
}

// Create: create an account in the database
func (r *AccountRepository) Create(account *Account) error {
	return r.db.Table(tbAccounts).Create(&account).Error
}

// FindOne: get a specific account inside the database
func (r *AccountRepository) FindOne(id int64) (*Account, error) {
	account := Account{}

	err := r.db.Table(tbAccounts).Model(&Account{}).Where("id = ?", id).Find(&account).Error

	return &account, err
}
