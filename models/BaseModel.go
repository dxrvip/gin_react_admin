package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type DeletedAt sql.NullTime

type BaseModel struct {
	ID        uint           `gorm:"primarykey;unique" json:"id"`
	CreatedAt time.Time      `gorm:"autoCreateTime:int" json:"createAd"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:int;<-:update" json:"updatedAd"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
