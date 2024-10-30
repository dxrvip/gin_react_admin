package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type JSONString []string

// Value 实现 driver.Valuer 接口
func (j JSONString) Value() (driver.Value, error) {
	return json.Marshal(j)
}

// Scan 实现 sql.Scanner 接口
func (j *JSONString) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, j)
}

type Role struct {
	BaseModel
	Name   string     `gorm:"type:string;size:50;not null;comment:名称" json:"name,omitempty"`
	Key    string     `gorm:"type:string;size:50;not null;comment:权限标识符" json:"key,omitempty"`
	Sort   uint       `gorm:"type:uint;default:0;comment:排序顺序" json:"sort,omitempty"`
	Active bool       `gorm:"type:bool;default:true;comment:是否启用" json:"active,omitempty"`
	Menus  JSONString `gorm:"type:text;comment:菜单" json:"menus,omitempty"`
}
