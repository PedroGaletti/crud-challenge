package operations

type Operation struct {
	ID   int64  `gorm:"column:id;primary_key;autoIncrement" json:"id"`
	Slug string `gorm:"column:slug;type:varchar(45);not null" json:"slug"`
}
