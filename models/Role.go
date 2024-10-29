package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Role struct {
	BaseModel
	Name   string   `gorm:"type:string;size:50;not null;comment:名称"`
	Key    string   `gorm:"type:string;size:50;not null;comment:权限标识符"`
	Sort   uint     `gorm:"type:uint;default:0;comment:排序顺序"`
	Active bool     `gorm:"type:bool;default:true;comment:是否启用"`
	Menu   []string `gorm:"type:json;comment:菜单"`
}

// Value实现数据库序列化的driver.Valuer接口
func (p Role) Value() (driver.Value, error) {
	return json.Marshal(p)
}

// Scan实现数据库反序列化的sql.Scanner接口
func (p *Role) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, p)
}
