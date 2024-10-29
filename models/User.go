package models

import (
	utils "goVueBlog/utils"

	"gorm.io/gorm"
)

type User struct {
	BaseModel
	Username string `gorm:"type:varchar(20);not null;index;unique;" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(500);not null" json:"password" validate:"required,min=6,max=120" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"`
	NikeName string `gorm:"type:varchar(50);default:nill" json:"nike_name" validate:"min=2,max=50" label:"昵称"`
	Email    string `gorm:"type:varchar(50);default:nill" json:"email" validate:"usage=email" label:"邮箱"`
	Active   bool   `gorm:"default:true" json:"status" validate:"required" label:"状态"`
	IsSuper  bool   `gorm:"default:false" label:"是否超级管理员"`
	Gender   string `gorm:"type:enum('male', 'female', 'other');default:'other'" json:"gender" label:"性别"`
}

// 插入之前进行加密
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password = utils.EncryptPassword(u.Password)
	return
}

// 获取密码
func (u *User) GetPassword() string {
	return "****"
}
