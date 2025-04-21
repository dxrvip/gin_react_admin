package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
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

type Float64String float64

func (fs *Float64String) UnmarshalJSON(data []byte) error {
	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch value := v.(type) {
	case float64:
		*fs = Float64String(value)
	case string:
		parsed, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return fmt.Errorf("invalid float string: %s", value)
		}
		*fs = Float64String(parsed)
	default:
		return fmt.Errorf("unsupported type: %T", v)
	}
	return nil

}
