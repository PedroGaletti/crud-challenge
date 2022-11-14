package transactions

import "time"

type Transaction struct {
	ID          int64     `gorm:"column:id;primary_key;autoIncrement" json:"id"`
	AccountID   int64     `gorm:"column:account_id;not null" json:"account_id" example:"321"`
	OperationID int64     `gorm:"column:operation_id;not null" json:"operation_id" example:"2"`
	Amount      int64     `gorm:"column:amount;not null" json:"amount" example:"2131"`
	CreatedAt   time.Time `gorm:"column:created_at;not null" json:"created_at"`
}
