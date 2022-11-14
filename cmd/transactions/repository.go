package transactions

import (
	"github.com/jinzhu/gorm"
)

const (
	tbTransactions = "transactions"
)

// ITransactionRepository: interface of Transaction repository
type ITransactionRepository interface {
	Create(*Transaction) error
}

// TransactionRepository: struct of Transaction repository
type TransactionRepository struct {
	db *gorm.DB
}

// NewTransactionRepository: create a new Transaction repository
func NewTransactionRepository(db *gorm.DB) ITransactionRepository {
	return &TransactionRepository{db}
}

// Create: create a transaction in the database
func (r *TransactionRepository) Create(transaction *Transaction) error {
	return r.db.Table(tbTransactions).Create(&transaction).Error
}
