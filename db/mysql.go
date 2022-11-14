package db

import (
	"fmt"

	"pismo/cmd/accounts"
	"pismo/cmd/operations"
	"pismo/cmd/transactions"
	"pismo/constants"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// InitMyqlDb: open MySQL database connection
func InitMyqlDb(user, password, host, port, database string) (*gorm.DB, error) {
	sn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=%s", user, password, host, port, database, "America%2FSao_Paulo")

	return gorm.Open("mysql", sn)
}

func Migrate(db *gorm.DB) {
	// tables
	db.AutoMigrate(&accounts.Account{}, &transactions.Transaction{}, &operations.Operation{})

	// constraints
	db.Model(&transactions.Transaction{}).AddForeignKey("account_id", "accounts(id)", "RESTRICT", "RESTRICT")
	db.Model(&transactions.Transaction{}).AddForeignKey("operation_id", "operations(id)", "RESTRICT", "RESTRICT")
}

func Seed(db *gorm.DB) {
	// operations types
	db.Model(&operations.Operation{}).Save((&operations.Operation{
		ID:   constants.Constants.Operations.Spot.ID,
		Slug: constants.Constants.Operations.Spot.Slug,
	}))
	db.Model(&operations.Operation{}).Save((&operations.Operation{
		ID:   constants.Constants.Operations.Installments.ID,
		Slug: constants.Constants.Operations.Installments.Slug,
	}))
	db.Model(&operations.Operation{}).Save((&operations.Operation{
		ID:   constants.Constants.Operations.Withdraw.ID,
		Slug: constants.Constants.Operations.Withdraw.Slug,
	}))
	db.Model(&operations.Operation{}).Save((&operations.Operation{
		ID:   constants.Constants.Operations.Payment.ID,
		Slug: constants.Constants.Operations.Payment.Slug,
	}))
}
