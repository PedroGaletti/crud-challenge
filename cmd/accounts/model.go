package accounts

type Account struct {
	ID             int64  `gorm:"column:id;primary_key;autoIncrement" json:"id"`
	DocumentNumber string `gorm:"column:document_number;not null" json:"document_number" example:"12345678900"`
}
