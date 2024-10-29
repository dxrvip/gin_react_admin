package models

import (
	"database/sql/driver"
	"encoding/json"

	"gorm.io/gorm"
)

type Role struct {
	BaseModel
	Name   string   `gorm:"type:string;size:50;not null;comment:名称"`
	Key    string   `gorm:"type:string;size:50;not null;comment:权限标识符"`
	Sort   uint     `gorm:"type:uint;default:0;comment:排序顺序"`
	Active bool     `gorm:"type:bool;default:true;comment:是否启用"`
	Menus  []string `gorm:"type:json;comment:菜单"`
}

// BeforeCreate 在插入之前序列化 Menus 字段
func (r *Role) Value(tx *gorm.DB) (driver.Value, error) {
	if len(r.Menus) > 0 {
		menusJSON, err := json.Marshal(r.Menus)
		if err != nil {
			return "[]", nil
		}
		return menusJSON, nil
	}
	return "[]", nil
}
