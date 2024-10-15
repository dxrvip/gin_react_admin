package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Picture struct {
	Src   string `gorm:"type:varchar(100)" json:"src" binding:"omitempty,url"`
	Title string `json:"title" binding:"max=100"`
}

// Value实现数据库序列化的driver.Valuer接口
func (p Picture) Value() (driver.Value, error) {
	return json.Marshal(p)
}

// Scan实现数据库反序列化的sql.Scanner接口
func (p *Picture) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, p)
}
