package models

type Role struct {
	BaseModel
	Name   string `gorm:"type:string;size:50;not null;comment:名称" json:"name,omitempty"`
	Key    string `gorm:"type:string;size:50;not null;comment:权限标识符" json:"key,omitempty"`
	Sort   uint   `gorm:"type:uint;default:0;comment:排序顺序" json:"sort,omitempty"`
	Active bool   `gorm:"type:bool;default:true;comment:是否启用" json:"active,omitempty"`
	Menus  string `gorm:"type:text;comment:菜单" json:"menus,omitempty"`
}

// func (m *Role) BeforeUpdate(tx *gorm.DB) (err error) {
// 	fmt.Println(tx, m)

// 	return
// }
