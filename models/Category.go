package models

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name" validate:"required,min(3),max(20)" label:"名称"`
}
