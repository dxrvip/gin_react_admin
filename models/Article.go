package models

import (
	"gorm.io/gorm"
)

type Article struct {
	Category Category `gorm:"foreignKey:Cid" json:"omitempty"`
	gorm.Model
	ID           uint    `gorm:"primarykey" json:"id"`
	Title        string  `gorm:"type:varchar(100);not null" json:"title"`
	Cid          int     `gorm:"type:int;not null" json:"cid"`
	Desc         string  `gorm:"type:varchar(200)" json:"desc"`
	Content      string  `gorm:"type:longtext" json:"content"`
	Picture      Picture `gorm:"type:json" json:"picture"`
	CommentCount int     `gorm:"type:int;not null;default:0" json:"comment_count"`
	ReadCount    int     `gorm:"type:int;not null;default:0" json:"read_count"`
}
